package routes

import (
	"asktils/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(500, gin.H{"message": "pong"})
	})

	router.GET("/api/clazz", handlers.GetAllClazzData(db))

	router.POST("/api/clazz/search", handlers.SerachClazzData(db))

	router.POST("/api/clazz/add", handlers.AddClazzData(db))

	router.POST("/api/clazz/delete", handlers.DeleteClazzData(db))

	router.POST("/api/clazz/edit", handlers.EditClazzData(db))
}
