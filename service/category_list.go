package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
	"go.uber.org/zap"
)

type CateListService struct{}

func (c *CateListService) CateList() (resp sz.Response) {
	dao := new(database.CateDao)
	data, err := dao.QueryAll()
	if err != nil {
		zap.S().Errorf("query failed: %e", err)
		return sz.Error()
	}
	resp = sz.Success()
	resp.Data = sz.NewMultiCategory(data)
	return
}
