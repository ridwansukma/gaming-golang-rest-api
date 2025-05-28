package main

import (
	"gamification/models"
	"gamification/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	if err := db.AutoMigrate(&models.GameScore{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// Seed data dummy (optional)
	seedDummyData(db)

	r := gin.Default()

	// Atur CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.1.107:3000"}, // sesuaikan dengan alamat frontend kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "ngrok-skip-browser-warning"},
		AllowCredentials: true,
	}))

	// Setup semua route di sini
	routes.SetupRoutes(r, db)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Merge Game API running")
	})

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}

func seedDummyData(db *gorm.DB) {
	// sama seperti sebelumnya
}
