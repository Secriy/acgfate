package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/util/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordPostService struct {
	Category int64  `json:"category" binding:"required,min=1,max=10"`
	Title    string `json:"title" binding:"required,min=1,max=48"`
	Content  string `json:"content" binding:"required,min=1,max=1024"`
}

func (service *WordPostService) Post(c *gin.Context) sz.Response {
	// get current user (author of word)
	user := model.CurrentUser(c)
	if user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr)
	}
	// check if the category exists.
	if cat, err := new(database.CatDao).QueryByID(service.Category); cat == nil && err == nil {
		// non-exists
		return sz.Failure()
	} else if err != nil {
		// other errors, maybe from database
		zap.S().Warnf("query error: %e", err)
		return sz.Error()
	}
	word := model.Word{
		Wid:     snowflake.Generate(),
		Aid:     user.UID,
		CatID:   service.Category,
		Title:   service.Title,
		Content: service.Content,
	}
	dao := new(database.WordDao)
	err := dao.Insert(&word)
	if err != nil {
		zap.S().Errorf("create word failed: %s", err)
		return sz.Error()
	}
	return sz.Success()
}
