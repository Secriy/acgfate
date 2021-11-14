package service

import (
	"acgfate/cache"
	"acgfate/database"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type WordDetailService struct{}

func (_ *WordDetailService) Detail(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	word, err := dao.QueryByID(c.Param("id"))
	if err != nil {
		return sz.Error()
	} else if word == nil {
		return sz.CodeResponse(sz.CodeWordNotExists)
	} else if word.IsDeleted() {
		return sz.CodeResponse(sz.CodeWordDeleted)
	}
	resp = sz.Success()
	word.Likes = new(cache.WordDao).Likes(c, word.Wid) // update data from cache
	resp.Data = sz.NewWord(word)
	return
}
