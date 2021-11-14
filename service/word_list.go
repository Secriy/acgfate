package service

import (
	"acgfate/cache"
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

const (
	DefaultPage = 1
	DefaultSize = 10
)

type WordListService struct {
	Page     int64 `json:"page" form:"page"`
	Size     int64 `json:"size" form:"size"`
	Author   int64 `json:"author" form:"author"`
	Category int64 `json:"category" form:"category"`
}

func (s *WordListService) List(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	// default paging
	if s.Page == 0 {
		s.Page = DefaultPage
	}
	if s.Size == 0 {
		s.Size = DefaultSize
	}
	// query by specific author or category,
	// if both exists, query using author.
	resp = sz.Success()
	if s.Author != 0 {
		// query by author
		words, err := dao.MQueryByAuthor(s.Author, s.Page-1, s.Size)
		if err != nil {
			return sz.Error()
		}
		updateLikes(c, words)
		resp.Data = sz.NewMultiWord(words)
		return
	}
	if s.Category != 0 {
		// query by category
		words, err := dao.MQueryByCat(s.Category, s.Page-1, s.Size)
		if err != nil {
			return sz.Error()
		}
		updateLikes(c, words)
		resp.Data = sz.NewMultiWord(words)
		return
	}
	// without specific author and category, query form all the data.
	words, err := dao.MQuery(s.Page-1, s.Size)
	if err != nil {
		return sz.Error()
	}
	updateLikes(c, words)
	resp.Data = sz.NewMultiWord(words)
	return
}

// updateLikes use data from cache to cover the old data.
func updateLikes(c *gin.Context, words []*model.Word) {
	for _, v := range words {
		v.UpdateLikes(new(cache.WordDao).Likes(c, v.Wid))
	}
}
