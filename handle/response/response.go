package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error, msg string, codes ...int) {
	code := http.StatusInternalServerError
	if len(codes) > 0 {
		code = codes[0]
	}

	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"err":  err.Error(),
	})
}
func Success[T any](c *gin.Context, data T, msg string) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}
