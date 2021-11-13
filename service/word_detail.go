package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordDetailService struct{}

func (_ *WordDetailService) Detail(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	wid := c.Param("id")
	word, err := dao.QueryByID(wid)
	if err != nil {
		zap.S().Warnf("get detail of %d error: %e", wid, err)
		return sz.Failure()
	}
	if word.Status == model.StatusWordDeleted {
		// word have been deleted
		return sz.CodeResponse(sz.CodeWordDeleted)
	}
	resp = sz.Success()
	resp.Data = sz.NewWord(word)
	return
}
