package v1

import (
	"net/http"

	"acgfate/service"
	"github.com/gin-gonic/gin"
)

// CategoryDetail 分区信息
func CategoryDetail(c *gin.Context) {
	serv := new(service.CatDetailService)
	res := serv.CatDetail(c.Param("name"))
	c.JSON(http.StatusOK, res)
}

// CategoryList 获取所有的分区列表
func CategoryList(c *gin.Context) {
	serv := new(service.CatListService)
	res := serv.CatList()
	c.JSON(http.StatusOK, res)
}
