package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
	"go.uber.org/zap"
)

type CatListService struct{}

func (c *CatListService) List() (resp sz.Response) {
	dao := new(database.CatDao)
	data, err := dao.QueryAll()
	if err != nil {
		zap.S().Errorf("query failed: %e", err)
		return sz.Error()
	}
	resp = sz.Success()
	resp.Data = sz.NewMultiCategory(data)
	return
}
