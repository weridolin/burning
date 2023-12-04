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
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  msg,
		"data": nil,
		"code": code,
	})
}
