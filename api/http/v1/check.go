package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/services"
	"github.com/gin-gonic/gin"
)

// CheckUsername 用户名接口
// @Summary 查询用户名是否存在
// @Description 查询用户名是否存在
// @Tags Check
// @Accept application/json
// @Produce  application/json
// @Param query query services.UsernameService true "用户名"
// @Success 0 {object} serializer.Response "msg: "可用"
// @Failure 40021 {object} serializer.Response "msg: 用户名已存在"
// @Router /check/username [get]
func CheckUsername(c *gin.Context) {
	var query services.UsernameService
	if err := c.ShouldBindQuery(&query); err == nil {
		res := query.CheckUsername()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParmErr())
	}
}

// CheckEmail 邮箱接口
// @Summary 查询邮箱是否存在
// @Description 查询邮箱是否存在
// @Tags Check
// @Accept application/json
// @Produce  application/json
// @Param query query services.EmailService true "邮箱"
// @Success 0 {object} serializer.Response "msg: "可用"
// @Failure 40022 {object} serializer.Response "msg: 邮箱已存在"
// @Router /user/email [get]
func CheckEmail(c *gin.Context) {
	var query services.EmailService
	if err := c.ShouldBindQuery(&query); err == nil {
		res := query.CheckEmail()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParmErr())
	}
}
