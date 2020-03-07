package main

import (
	"github.com/alitaso345/azusaconverter"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type postRequest struct {
	Content string `json:content`
}

// 佐々木梓の台詞には「[azusa:久美子さ、なんで南宇治高校に行かんかったん？]」のようなアノテーションがついている前提
func convertAnnotation(text string) string {
	converted := ""
	for _, line := range strings.Split(text, "\n") {
		r := regexp.MustCompile(`[azusa:*]`)
		if r.MatchString(line) {
			line = strings.Replace(line, "[azusa:", "", 1)
			line = strings.Replace(line, "]", "", 1)
			converted += azusaconverter.Convert(line)
		} else {
			converted += line
		}
		converted += "\n"
	}
	return converted
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.tmpl.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	r.POST("/convert", func(c *gin.Context) {
		requestBody := postRequest{}
		c.Bind(&requestBody)

		result := convertAnnotation(requestBody.Content)
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
