package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryPath r423423
func QueryPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})

	}
}
