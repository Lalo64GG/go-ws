package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)


func ConfigurationCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500"},
        AllowMethods:      []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:      []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:      []string{"Content-Length"},
        AllowCredentials:  true,
        MaxAge:            12 * time.Hour,
	})
}