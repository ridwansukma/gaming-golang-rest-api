package routes

import (
	"gamification/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.DB = db

	api := r.Group("/api")
	{
		api.POST("/submit-score", controllers.SubmitScore)
		api.GET("/leaderboard", controllers.GetLeaderboard)
		api.DELETE("/leaderboard", controllers.ClearLeaderboard)

	}
}
