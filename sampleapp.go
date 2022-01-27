package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		response := fmt.Sprintf("%s's email is %s", name, email)

		c.String(http.StatusOK, response)
	})

	router.Run(":8080")
}
