package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"golang-learning/practise/gin-practise/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(middleware.LoggerToFile())

	r.GET("/testlog", middleware.Valid, TestLog)

	r.Run(":8000")
}

func TestLog(c *gin.Context) {
	fmt.Println("testlog")
	c.JSON(200, map[string]string{"test": "log"})
	return
}
