package middleware

import (
	"fmt"
	"simple-block-api/models"
	"simple-block-api/utils"

	"github.com/gin-gonic/gin"
)

func MyRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				switch e := err.(type) {
				case string:
					errMsg = e
				case error:
					errMsg = e.Error()
				default:
					errMsg = "Unknown error: " + fmt.Sprintf("%v", e)
				}

				// log error to database:
				db := utils.GetDB()
				db.Save(&models.ErrorLog{Message: errMsg})

				utils.RespWithHttpError(c, 500, errMsg)
				c.Abort()
			}
		}()
		c.Next()
	}
}
