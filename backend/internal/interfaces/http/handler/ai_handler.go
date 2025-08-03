package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/ai"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type AiHandler struct {
	cnfg config.Config
}

func NewAiHandler(cnfg config.Config) *AiHandler {
	return &AiHandler{
		cnfg: cnfg,
	}
}

func (h *AiHandler) CalculationCalories(c *gin.Context) {
	var cReq dto.CaloriesRequest

	if err := c.ShouldBindJSON(&cReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	jsonData, err := ai.GenerateMessage(
		string(h.cnfg.CaloriesAnalyzer),
		fmt.Sprintf("%s %s", h.cnfg.FoodToAnalyze, cReq.Title))

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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var aiResp ai.APIResponse

	if err := ai.Deserialization(bodyBytes, &aiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	content := strings.TrimSpace(aiResp.Choices[0].Message.Content)

	fmt.Println("AI raw content:", content)

	if content == "null" || content == "\"null\"" {
		c.JSON(http.StatusOK, dto.CaloriesResponse{Calories: nil})
		return
	}

	content = strings.Trim(content, "\"")

	calories, err := strconv.ParseFloat(content, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to parse calories float",
			"content": content,
		})
		return
	}
	caloriesRounded := int(math.Ceil(calories))

	c.JSON(http.StatusOK, dto.CaloriesResponse{Calories: &caloriesRounded})
}
