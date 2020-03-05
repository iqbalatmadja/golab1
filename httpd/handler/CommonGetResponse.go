package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Template1 r423423
func Template1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "titling",
		})
	}
}
