package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
		"code": code,
	})
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"msg":  msg,
		"data": nil,
		"code": code,
	})
}
