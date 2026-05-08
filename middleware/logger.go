package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyLogger(c *gin.Context) {
	fmt.Printf("Request 【%s】start...\n", c.Request.URL.Path)
	c.Next()
	fmt.Printf("Request 【%s】end\n\n", c.Request.URL.Path)
}
