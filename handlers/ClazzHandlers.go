package handlers

import (
	"asktils/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllClazzData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clazzDatas []models.ClazzData
		result := db.Find(&clazzDatas)
		if result.Error != nil {
			c.JSON(200, gin.H{"error": "查询失败"})
			return
		}
		c.HTML(200, "home/index.html", gin.H{"Data": clazzDatas})
	}
}
