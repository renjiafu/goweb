package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mes": time.Now(),
		})
	})

	router.GET("/index1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mes": "Hello go web!",
		})
	})

	router.GET("/index2", func(c *gin.Context) {
		var msg struct {
			Name string
			Mes  string
			Age  int
		}
		msg.Name = "小王子"
		msg.Mes = "Hello GO WEB!"
		msg.Age = 23
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	router.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	router.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	router.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	router.Run()

}
