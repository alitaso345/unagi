package main

import (
	"net/http"
	"os"
)
import "github.com/gin-gonic/gin"

func main() {
	port := os.Getenv("PORT")

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "佐々木梓さんお誕生日おめでとうございます",
		})
	})

	r.Run(":" + port)
}
