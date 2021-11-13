package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
	"go.uber.org/zap"
)

type WordDetailService struct{}

func (w *WordDetailService) Detail(wid interface{}) (resp sz.Response) {
	dao := new(database.WordDao)
	word, err := dao.QueryByID(wid)
	if err != nil {
		zap.S().Warnf("get detail of %d error: %e", wid, err)
		return sz.Failure()
	}
	resp = sz.Success()
	resp.Data = sz.NewWord(word)
	return
}
