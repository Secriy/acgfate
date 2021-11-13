package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/service"
	"github.com/gin-gonic/gin"
)

// CategoryDetail 分区信息
func CategoryDetail(c *gin.Context) {
	form := new(service.CateDetailService)
	if err := c.ShouldBind(form); err == nil {
		res := form.CateDetail()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParamErrorResponse())
	}
}

// CategoryList 获取所有的分区列表
func CategoryList(c *gin.Context) {
	form := new(service.CateListService)
	if err := c.ShouldBind(form); err == nil {
		res := form.CateList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParamErrorResponse())
	}
}
