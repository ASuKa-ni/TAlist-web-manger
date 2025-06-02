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
			c.HTML(500, "home/error.html", nil)
			return
		}
		c.HTML(200, "home/index.html", gin.H{"Data": clazzDatas})
	}
}

func SerachClazzData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchCriteria models.ClazzData
		var clazzDatas []models.ClazzData

		if err := c.ShouldBind(&searchCriteria); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request data"})
			return
		}

		query := db.Model(&models.ClazzData{})

		if searchCriteria.Id != 0 {
			query = query.Where("id = ?", searchCriteria.Id)
		}
		if searchCriteria.DataClazzName != "" {
			query = query.Where("data_clazz_name = ?", searchCriteria.DataClazzName)
		}
		if searchCriteria.DataClazzRoom != "" {
			query = query.Where("data_clazz_room = ?", searchCriteria.DataClazzRoom)
		}
		if searchCriteria.DataClazzTeacher != "" {
			query = query.Where("data_clazz_teacher = ?", searchCriteria.DataClazzTeacher)
		}
		if searchCriteria.DataClazzStatus != "" {
			query = query.Where("data_clazz_status = ?", searchCriteria.DataClazzStatus)
		}

		result := query.Find(&clazzDatas)
		if result.Error != nil {
			c.HTML(500, "home/error.html", nil)
			return
		}

		c.HTML(200, "home/index.html", gin.H{"Data": clazzDatas})
	}
}

func AddClazzData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clazzData models.ClazzData

		if err := c.ShouldBind(&clazzData); err != nil {
			c.AbortWithStatus(400)
			return
		}

		if clazzData.DataClazzName == "" {
			c.HTML(200, "home/message.html", gin.H{"msg": "不能为空"})
			return
		}
		if clazzData.DataClazzRoom == "" {
			c.HTML(200, "home/message.html", gin.H{"msg": "不能为空"})
			return
		}
		if clazzData.DataClazzTeacher == "" {
			c.HTML(200, "home/message.html", gin.H{"msg": "不能为空"})
			return
		}
		if clazzData.DataClazzStatus == "" {
			c.HTML(200, "home/message.html", gin.H{"msg": "不能为空"})
			return
		}

		result := db.Create(&clazzData)
		if result.Error != nil {
			c.HTML(400, "home /message.html", gin.H{"msg": "错误"})
			return
		}
		c.HTML(200, "home/message.html", gin.H{"msg": "添加成功！"})
	}
}

func DeleteClazzData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clazzData models.ClazzData

		if err := c.ShouldBind(&clazzData); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request data"})
			return
		}

		if clazzData.Id == 0 {
			c.HTML(200, "home/message.html", gin.H{"msg": "ID不能为空"})
			return
		}

		result := db.Delete(&clazzData)
		if result.Error != nil {
			c.HTML(500, "home/error.html", nil)
			return
		}

		c.HTML(200, "home/message.html", gin.H{"msg": "删除成功!"})
	}
}

func EditClazzData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clazzData models.ClazzData

		if err := c.ShouldBind(&clazzData); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request data"})
			return
		}

		if clazzData.Id == 0 {
			c.HTML(200, "home/message.html", gin.H{"msg": "ID不能为空"})
			return
		}

		result := db.Model(&models.ClazzData{}).Where("id = ?", clazzData.Id).Updates(map[string]interface{}{
			"data_clazz_name":    clazzData.DataClazzName,
			"data_clazz_room":    clazzData.DataClazzRoom,
			"data_clazz_teacher": clazzData.DataClazzTeacher,
			"data_clazz_status":  clazzData.DataClazzStatus,
		})

		if result.Error != nil {
			c.HTML(500, "home/error.html", nil)
			return
		}

		c.HTML(200, "home/message.html", gin.H{"msg": "修改成功!"})
	}
}
