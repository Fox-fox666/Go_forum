package controllers

import (
	"Go_forum/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = 666
	nono    = 2333
)

func ResErr(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, models.Response{
		Code: nono,
		Msg:  msg,
		Data: data,
	})
}

func ResSuccess(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, models.Response{
		Code: success,
		Msg:  msg,
		Data: data,
	})
}
