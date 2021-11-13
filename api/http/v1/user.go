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
// @Produce  application/json
// @Param form body service.UserRegisterService true "用户名, 密码, 邮箱, 昵称"
// @Success 0 {object} serializer.Response "msg: Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
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
// @Summary 用户登录
// @Description 用户登录接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param form body service.UserLoginService true "用户名, 密码"
// @Success 0 {object} serializer.Response "msg: Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
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
// @Summary 用户个人信息
// @Description 用户个人信息接口
// @Tags User
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Success 0 {object} serializer.User "msg: Success"
// @Failure 50000 {object} serializer.Response "msg: 查询个人信息错误"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	form := new(service.UserInfoService)
	res := form.Info(c)
	c.JSON(http.StatusOK, res)
}

// UserInfoUpdate 用户信息更新
// @Summary 用户信息更新
// @Description 用户信息更新接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Param form body service.UpdateInfoService true "用户信息"
// @Success 0 {object} serializer.BasicInfoResponse "msg: "Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /user/update [put]
// func UserInfoUpdate(c *gin.Context) {
// 	var form service.UpdateInfoService
// 	if err := c.ShouldBind(&form); err == nil {
// 		res := form.Update(c)
// 		c.JSON(http.StatusOK, res)
// 	} else {
// 		c.JSON(http.StatusOK, sz.ParamErr())
// 	}
// }
