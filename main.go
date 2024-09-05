package main

import (
	"github.com/ggorockee/webserver-with-go/controllers"
	"github.com/ggorockee/webserver-with-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	models.ConnectDB()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", controllers.CreateUser)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
