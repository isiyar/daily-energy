package handler

import (
	"encoding/json"
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
		"You are a calorie-estimation assistant.\nWhen given the name of a food or dish, you must:\n\nIdentify the typical serving size based on general knowledge (e.g., restaurant portion or common packaging).\n\nEstimate the number of kilocalories (Calories) in a typical serving.\n\nIf no amount is specified, assume a standard portion.\n\n️ You must respond with only one float number, rounded to one decimal place.\nDo not add any explanation, description, or text — only output the number itself.\n\n YES Example 1:\nInput: \"Boiled egg\"\nOutput:\n78.0\n\nYES Example 2:\nInput: \"Chocolate bar\"\nOutput:\n230.0\n\nNOT Incorrect:\n\"A chocolate bar contains around 230 calories.\" ← This is forbidden.\n\nIf the item is truly unidentifiable, return:\nnull",
		fmt.Sprintf("Сколько калорий в стандартной порции %s", cReq.Title))

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

	//c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var aiResp dto.CaloriesAPIResponse
	if err := json.Unmarshal(bodyBytes, &aiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse AI response"})
		return
	}

	if len(aiResp.Choices) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No choices in AI response"})
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
