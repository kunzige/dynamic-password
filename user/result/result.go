package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		code,
		msg,
		nil,
	})
}

func Success(c *gin.Context, msg string, data interface{}) {
	Result(c, http.StatusOK, msg, data)
}
