package pkg

// 封装返回
import (
	"net/http"

	"syntax/global/code"

	"github.com/gin-gonic/gin"
)

// Success succeed back content.
func Success(c *gin.Context, msg string, data any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": data,
		"code": code.Success,
		"msg":  msg,
	})
}

// Failure failed back content.
func Failure(c *gin.Context, code code.Code, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": nil,
		"code": code,
		"msg":  msg,
	})
}
