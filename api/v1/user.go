package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/services/user"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// @Summary 用户注册
// @Description 用户注册接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param form body user.RegisterService true "用户名, 密码, 邮箱, 昵称"
// @Success 0 {object} serializer.UserResponse "msg: Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var form user.RegisterService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ErrorMsg(sz.ParamErr, err))
	}
}

// UserLogin 用户登录接口
// @Summary 用户登录
// @Description 用户登录接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param form body user.LoginService true "用户名, 密码"
// @Success 0 {object} serializer.LoginResponse "msg: Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var form user.LoginService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ErrorMsg(sz.ParamErr, err))
	}
}

// UserMe 用户个人信息接口
// @Summary 用户个人信息
// @Description 用户个人信息接口
// @Tags User
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Success 0 {object} serializer.UserResponse "msg: Success"
// @Failure 50001 {object} serializer.Response "msg: 查询个人信息错误"
// @Router /user/me [get]
func UserMe(c *gin.Context) {
	var form user.MeService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Me(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ErrorMsg(sz.ParamErr, err))
	}
}

// UserUpdate 用户信息更新
// @Summary 用户信息更新
// @Description 用户信息更新接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Param form body user.UpdateService true "用户信息"
// @Success 0 {object} serializer.UserResponse "msg: "Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /user/update [post]
func UserUpdate(c *gin.Context) {
	var form user.UpdateService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Update(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ErrorMsg(sz.ParamErr, err))
	}
}
