package controller

import (
	"errors"
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler 注册
// @Summary 注册接口
// @Description 注册
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamSignUp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} RespData
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 请求参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			RespError(c, CodeInvalidParam)
			return
		}
		RespErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 手动对请求参数进行详细的业务规则校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	// 	zap.L().Error("SignUp with invalid param", zap.Error(err))
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"meg": "请求参数有误",
	// 	})
	// 	return
	// }

	fmt.Printf("p: %v\n", p)

	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("service.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			RespError(c, CodeUserExist)
			return
		}
		RespError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	RespSuccess(c, nil)
}

// LoginHandler 登录
// @Summary 登录接口
// @Description 登录
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamLogin false "查询参数"
// @Success 200 {object} RespData
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数及校验
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 请求参数有误
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			RespError(c, CodeInvalidParam)
			return
		}
		RespErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2.业务处理逻辑
	user, err := service.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			RespError(c, CodeUserNotExist)
		}
		RespError(c, CodeInvalidPassword)
		return
	}

	// 3.返回响应
	RespSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // 把id转成字符串
		"user_name": user.Username,
		"token":     user.Token,
	})
}
