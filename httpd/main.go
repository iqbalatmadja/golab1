package main

import (
	"fmt"
	"golab1/httpd/handler"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping fwefde
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// PostResponse 423432
func PostResponse(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"body": string(value),
	})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", Ping)
	r.GET("/test/:id/:name", handler.QueryPath())
	r.POST("/test", PostResponse)
	r.GET("/template1", handler.Template1())

	r.Run(":9999")
}
