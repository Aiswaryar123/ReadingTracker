package main

import (
	"readingtracker/configs"
	"readingtracker/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	
	database.ConnectDB(cfg)

	
	r := gin.Default()



	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	
	r.Run(":" + cfg.Port)
}
