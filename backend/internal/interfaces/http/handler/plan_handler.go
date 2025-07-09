package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/ai"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type PlanHandler struct {
	cnfg   config.Config
	planUC *usecase.PlanUseCase
	userUC *usecase.UserUseCase
}

func NewPlanHandler(cnfg config.Config, planUC *usecase.PlanUseCase, userUC *usecase.UserUseCase) *PlanHandler {
	return &PlanHandler{
		cnfg:   cnfg,
		planUC: planUC,
		userUC: userUC,
	}
}

func (h *PlanHandler) CreatePlan(c *gin.Context) {
	var Preq dto.PlanRequest

	if err := c.ShouldBindJSON(&Preq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := validator.Struct(Preq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	utgidStr := c.Param("utgid")

	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	user, err := h.userUC.Execute(c.Request.Context(), utgid)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(fmt.Sprintf("%s %d %d %d %s %s %d", user.Gender, user.DateofBirth, user.Weight, user.Height, user.Goal, user.PhysicalActivity, Preq.Date))

	jsonData, err := ai.GenerateMessage(
		ai.PlanGenerator,
		fmt.Sprintf("%s %d %d %d %s %s %d", user.Gender, user.DateofBirth, user.Weight, user.Height, user.Goal, user.PhysicalActivity, Preq.Date))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode request body"})
		return
	}

	req, err := ai.GenerateRequest(h.cnfg, jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	var aiResp ai.APIResponse

	if err := ai.Deserialization(bodyBytes, &aiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	content := strings.TrimSpace(aiResp.Choices[0].Message.Content)

	fmt.Println("AI raw content:", content)

	content = strings.Trim(content, "\"")

	fmt.Println(content)

	var planContent dto.AIPlanContent
	if err := json.Unmarshal([]byte(content), &planContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse AI plan content", "details": err.Error()})
		return
	}

	var plans []models.Plan

	for dateStr, nutrition := range planContent.Nutrition {
		date, err := strconv.ParseInt(dateStr, 10, 64)
		if err != nil {
			continue
		}

		plan := models.Plan{
			Id:                uuid.New().String(),
			Utgid:             utgid,
			Date:              date,
			CaloriesToConsume: nutrition.Calories,
			CaloriesToBurn:    0,
			Recommendation:    strings.Join(nutrition.Recommendations, "\n"),
			Type:              models.Food,
		}

		plans = append(plans, plan)
	}

	for dateStr, workout := range planContent.Workouts {
		date, err := strconv.ParseInt(dateStr, 10, 64)
		if err != nil {
			continue
		}

		plan := models.Plan{
			Id:                uuid.New().String(),
			Utgid:             utgid,
			Date:              date,
			CaloriesToConsume: 0,
			CaloriesToBurn:    workout.Calories,
			Recommendation:    strings.Join(workout.Recommendations, "\n"),
			Type:              models.Activity,
		}

		plans = append(plans, plan)
	}

	err = h.planUC.Add(c, plans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse AI plan content", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToPlansResponse(plans))

	//c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

func (h *PlanHandler) GetPlans(c *gin.Context) {

}
