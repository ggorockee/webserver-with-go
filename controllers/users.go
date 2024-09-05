package controllers

//
//import (
//	"github.com/ggorockee/webserver-with-go/models"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type CreateUserSerializer struct {
//	Name     string `json:"name" binding:"required"`
//	Email    string `json:"email" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}
//
//type TinyUserSerializer struct {
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
//
//type UpdateUserSerializer struct {
//	Name string `json:"name"`
//}
//
//func CreateUser(c *gin.Context) {
//	input := CreateUserSerializer{}
//	if err := c.ShouldBind(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	user := models.User{
//		Name:     input.Name,
//		Email:    input.Email,
//		Password: input.Password,
//	}
//	user.SetPassword()
//	result := models.Database.DB.Create(&user)
//	if result.Error != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	// serializer User
//	responseUser := TinyUserSerializer{
//		Name:  user.Name,
//		Email: user.Email,
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"status": "ok",
//		"data":   responseUser,
//	})
//}
