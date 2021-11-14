package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type WordDeleteService struct{}

func (_ *WordDeleteService) Delete(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	word, err := dao.QueryByID(c.Param("id"))
	if err != nil {
		return sz.Error()
	} else if word == nil {
		return sz.CodeResponse(sz.CodeWordNotExists)
	} else if word.IsDeleted() {
		return sz.CodeResponse(sz.CodeWordDeleted)
	}

	// get current user (author of word)
	if user := model.CurrentUser(c); user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr)
	} else if user.UID != word.Aid {
		// author not match
		return sz.Failure()
	}

	word.Status = model.StatusWordDeleted // set status to deleted

	if err := dao.Update(word); err != nil {
		return sz.Error()
	}
	resp = sz.Success()
	return
}
