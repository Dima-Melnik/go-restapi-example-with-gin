package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Type   string
	Method string
	Resp   interface{}
}

func SendErrorResp(c *gin.Context, status int, payload interface{}) {
	resp := Response{
		Type:   "Error",
		Method: c.Request.Method,
		Resp:   payload,
	}

	c.JSON(status, resp)
	c.Abort()
}

func SendJsonResp(c *gin.Context, status int, payload interface{}) {
	resp := Response{
		Type:   "Success",
		Method: c.Request.Method,
		Resp:   payload,
	}

	c.JSON(status, resp)
}
