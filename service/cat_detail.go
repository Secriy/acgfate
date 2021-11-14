package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CatDetailService struct{}

// Detail return the response of category.
func (_ *CatDetailService) Detail(c *gin.Context) (resp sz.Response) {
	dao := new(database.CatDao)
	name := c.Param("name")
	cat, err := dao.QueryByCname(name)
	if err != nil {
		zap.S().Warnf("query name %s error: %e", name, err)
		return sz.Failure()
	}
	resp = sz.Success()
	resp.Data = sz.NewCategory(cat)
	return
}
