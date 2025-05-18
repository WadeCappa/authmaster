package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const AUTH_PREFIX = "Authorization"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		token := c.Request.Header[AUTH_PREFIX]
		c.JSON(http.StatusOK, gin.H{
			"token_data": c.Request.Header["Authorization"],
		})
	})

	return r
}

func getToken(c *gin.Context) string {
	return c.Request.Header[AUTH_PREFIX]
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
