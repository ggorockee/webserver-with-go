package controllers

import (
	"github.com/ggorockee/webserver-with-go/database"
	"github.com/ggorockee/webserver-with-go/models"
	"github.com/ggorockee/webserver-with-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound database.User
	database.Database.DB.Where("email = ?", authInput.Email).Find(&userFound)

	if userFound.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := utils.CheckPassword(userFound.Password, authInput.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

}
