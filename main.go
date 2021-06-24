package main

import "github.com/gin-gonic/gin"

func createEntity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func updateEntity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func deleteEntity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.PUT("/create", createEntity)
	r.POST("/update", updateEntity)
	r.DELETE("/delete", deleteEntity)
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
