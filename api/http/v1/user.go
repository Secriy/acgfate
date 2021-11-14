package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/service"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// @Summary 用户注册
// @Description 用户注册接口
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param register body service.UserRegisterService true "用户名, 密码, 邮箱, 昵称"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	form := new(service.UserRegisterService)
	if err := c.ShouldBind(form); err == nil {
		res := form.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParamError())
	}
}

// UserLogin 用户登录接口
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param login body service.UserLoginService true "用户名, 密码"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	form := new(service.UserLoginService)
	if err := c.ShouldBind(form); err == nil {
		res := form.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParamError())
	}
}

// UserInfo 用户个人信息接口
// @Summary 用户个人信息接口
// @Description 用户个人信息，要求登录
// @Tags User
// @Produce application/json
// @Success 0 {object} serializer.User "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	form := new(service.UserInfoService)
	res := form.Info(c)
	c.JSON(http.StatusOK, res)
}
