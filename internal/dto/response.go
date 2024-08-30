package dto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var zapLogger, _ = zap.NewDevelopment()

type SuccessResponse struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Data: data,
		Code: http.StatusOK,
	})
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErrorResponse(c *gin.Context, code int, err error, message string) {
	if err == nil {
		err = errors.New(message)
	}
	zapLogger.Error(fmt.Sprintf("Error: %s", err.Error()))
	c.AbortWithStatusJSON(code, ErrorResponse{
		Error:   err.Error(),
		Message: message,
		Code:    code,
	})
}
