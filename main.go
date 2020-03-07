package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.tmpl.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	r.POST("/convert", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "佐々木梓さんお誕生日おめでとうございます",
		})
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
