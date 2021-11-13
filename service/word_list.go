package service

import (
	"strconv"

	"acgfate/database"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordListService struct{}

func (_ *WordListService) List(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	// paging
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 1 // default offset
	}
	number, err := strconv.ParseInt(c.Query("number"), 10, 64)
	if err != nil {
		number = 10 // default limit
	}
	// query by specific author or category,
	// if both exists, query using author.
	author, _ := strconv.ParseInt(c.Query("author"), 10, 64)
	category, _ := strconv.ParseInt(c.Query("category"), 10, 64)

	resp = sz.Success()
	if author != 0 {
		// query by author
		zap.S().Info(page)
		words, err := dao.MQueryByAuthor(author, page-1, number)
		if err != nil {
			return sz.Error()
		}
		resp.Data = sz.NewMultiWord(words)
		return
	}
	if category != 0 {
		// query by category
		words, err := dao.MQueryByCat(category, page-1, number)
		if err != nil {
			return sz.Error()
		}
		resp.Data = sz.NewMultiWord(words)
		return
	}
	// without specific author and category, query form all the data.
	words, err := dao.MQuery(page-1, number)
	if err != nil {
		return sz.Error()
	}
	resp.Data = sz.NewMultiWord(words)
	return
}
