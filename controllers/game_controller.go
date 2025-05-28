package controllers

import (
	"gamification/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScoreInput struct {
	Username string `json:"username"`
	Merges   int    `json:"merges"`
	Duration int    `json:"duration"`
}

var DB *gorm.DB // akan diassign dari main.go

func SubmitScore(c *gin.Context) {
	var input ScoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	reward := 0
	if input.Merges >= 10 && input.Duration <= 60 {
		reward = 20
	}

	score := models.GameScore{
		Username: input.Username,
		Merges:   input.Merges,
		Duration: input.Duration,
		Reward:   reward,
	}

	if err := DB.Create(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save score"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reward": reward})
}

func GetLeaderboard(c *gin.Context) {
	var leaderboard []models.GameScore

	if err := DB.Order("reward DESC, duration ASC").Limit(10).Find(&leaderboard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if leaderboard == nil {
		leaderboard = []models.GameScore{}
	}

	c.JSON(http.StatusOK, leaderboard)
}

func ClearLeaderboard(c *gin.Context) {
	if err := DB.Where("1 = 1").Delete(&models.GameScore{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear leaderboard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leaderboard cleared"})
}
