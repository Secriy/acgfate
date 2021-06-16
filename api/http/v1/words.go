package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/services"
	"github.com/gin-gonic/gin"
)

// WordsPost 发表
// @Summary 文字发表
// @Description 文字发表接口
// @Tags Words
// @Accept application/json
// @Produce  application/json
// @Param form body services.PostService true "文字发表信息"
// @Success 0 {object} serializer.Response "msg:"发表成功""
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /words/post [post]
func WordsPost(c *gin.Context) {
	var form services.PostService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Post(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParmErr())
	}
}

// WordsGet 获取
// @Summary 文字查看
// @Description 文字查看接口
// @Tags Words
// @Produce  application/json
// @Param WID query string true "Words ID"
// @Success 0 {object} serializer.WordsResponse "msg:"获取成功""
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /words/:id [get]
func WordsGet(c *gin.Context) {
	var form services.GetService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Get(c.Param("wid"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParmErr())
	}
}
