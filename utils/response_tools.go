package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func RespWithMessage(ctx *gin.Context, message string) {
	CommonResponse(ctx, http.StatusOK, 0, message, nil)
}

func RespWithMessageAndData(ctx *gin.Context, message string, data any) {
	CommonResponse(ctx, http.StatusOK, 0, message, data)
}

func RespWithData(ctx *gin.Context, data any) {
	CommonResponse(ctx, http.StatusOK, 0, "Success", data)
}

func RespWithBizError(ctx *gin.Context, errorCode *ErrorCode) {
	CommonResponse(ctx, http.StatusOK, errorCode.Code, errorCode.Message, nil)
}

func RespWithHttpError(ctx *gin.Context, httpCode int, message string) {
	CommonResponse(ctx, httpCode, 0, message, nil)
}

func CommonResponse(ctx *gin.Context, httpCode int, errorCode int, message string, data any) {
	ctx.JSON(httpCode, ResponseModel{
		Code:    errorCode,
		Message: message,
		Data:    data,
	})
}
