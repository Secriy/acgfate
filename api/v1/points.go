package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/services/points"
	"github.com/gin-gonic/gin"
)

// SignService 用户信息更新
// @Summary 用户信息更新
// @Description 用户信息更新接口
// @Tags User
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Param form body points.SignService true "用户信息"
// @Success 0 {object} serializer.UserPointsResponse "msg: "Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /sign [post]
func SignService(c *gin.Context) {
	var form points.SignService
	if err := c.ShouldBind(&form); err == nil {
		res := form.DoSign(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ErrorResponse(
			sz.ParamErr, sz.GetResMsg(sz.ParamErr), err))
	}
}
