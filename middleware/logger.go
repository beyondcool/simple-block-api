package middleware

import "github.com/gin-gonic/gin"

func MyLogger(c *gin.Context) {

	c.Next()

}
