package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lzh-tsak/model"
)

func render(c *gin.Context,data interface{},err error) {
	response := model.Response{
		Code:    200,
		Success: func() bool{
			if err != nil {
				return false
			}
			return true
		}(),
		Data: func() interface{}{
			if err != nil {
				return nil
			}
			return data
		}(),
	}
	content,err := json.Marshal(response)
	if err != nil {
		c.Abort()
	}
	c.Writer.Write(content)
}
