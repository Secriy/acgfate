package v1

import (
	"net/http"

	"acgfate/service"
	"github.com/gin-gonic/gin"
)

// CategoryDetail 分区信息接口
// @Summary 分区信息接口
// @Description 分区信息接口
// @Tags Category
// @Produce application/json
// @Param name path string true "分区名称"
// @Success 0 {object} serializer.Category "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /category/{name} [get]
func CategoryDetail(c *gin.Context) {
	serv := new(service.CatDetailService)
	res := serv.Detail(c)
	c.JSON(http.StatusOK, res)
}

// CategoryList 获取所有的分区列表
// @Summary 分区信息接口
// @Description 分区信息接口
// @Tags Category
// @Produce application/json
// @Success 0 {array} []serializer.Category "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /category/list [get]
func CategoryList(c *gin.Context) {
	serv := new(service.CatListService)
	res := serv.List()
	c.JSON(http.StatusOK, res)
}
