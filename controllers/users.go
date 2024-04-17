package controllers

import (
	"Go_forum/common"
	"Go_forum/logic"
	"Go_forum/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 注册
func RegisterHandler(c *gin.Context) {
	//1、参数校验
	register := new(models.Register)

	//校验
	err := c.ShouldBindJSON(register)
	if err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Register with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResErr(c, err.Error(), nil)
			return
		}
		ResErr(c, "", common.RemoveTopStruct(errs.Translate(common.Trans)))
		return
	}

	//2、业务处理
	err = logic.Register(register)
	if err != nil {
		zap.L().Error("logic.Register failed:", zap.Error(err))
		ResErr(c, err.Error(), nil)
		return
	}
	//3、返回响应
	ResSuccess(c, "Register success", nil)
}

// 登录
func Login(c *gin.Context) {
	//获取请求参数并校验
	l := new(models.Login)
	err := c.ShouldBindJSON(l)
	if err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResErr(c, err.Error(), nil)
			return
		}
		ResErr(c, "", common.RemoveTopStruct(errs.Translate(common.Trans)))
		return
	}
	//业务逻辑处理
	token, err := logic.Login(l)
	if err != nil {
		zap.L().Error("logic.Login failed:", zap.String("username", l.Username), zap.Error(err))
		ResErr(c, err.Error(), nil)
		return
	}

	//返回响应
	ResSuccess(c, "登录成功", token)
}
