package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
	"go.uber.org/zap"
)

type CatDetailService struct{}

// Detail return the response of category.
func (c *CatDetailService) Detail(catName string) (resp sz.Response) {
	dao := new(database.CatDao)
	cat, err := dao.QueryByCname(catName)
	if err != nil {
		zap.S().Warnf("query name %s error: %e", catName, err)
		return sz.Failure()
	}
	resp = sz.Success()
	resp.Data = sz.NewCategory(cat)
	return
}
