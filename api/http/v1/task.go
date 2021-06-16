package v1

import (
	"net/http"

	"acgfate/services"
	"github.com/gin-gonic/gin"
)

// TaskSign 普通签到
// @Summary 普通签到
// @Description 普通签到接口
// @Tags Task
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Success 0 {object} serializer.TaskSignResponse "msg: "Success"
// @Failure 60001 {object} serializer.Response "msg: 参数错误"
// @Router /sign [post]
func TaskSign(c *gin.Context) {
	var form services.SignService
	res := form.DoSign(c)
	c.JSON(http.StatusOK, res)
}
