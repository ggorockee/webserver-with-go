package controllers

import (
	"github.com/ggorockee/webserver-with-go/models"
	"github.com/ggorockee/webserver-with-go/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateMemo(c *gin.Context) {
	payload := serializers.CreateMemoSerializer{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memo := models.Memo{
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := models.Database.DB.Create(&memo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "crated",
		"data":   nil,
	})
}
