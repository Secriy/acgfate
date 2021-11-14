package service

import (
	"acgfate/cache"
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordLikeService struct{}

func (_ *WordLikeService) Like(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	wid := c.Param("id")
	if word, err := dao.QueryByID(wid); err != nil {
		zap.S().Warnf("query word error: %e", err)
		return sz.Error()
	} else if word == nil {
		return sz.Failure() // word not exists
	} else if word.IsDeleted() {
		// word have been deleted
		return sz.CodeResponse(sz.CodeWordDeleted)
	}
	user := model.CurrentUser(c)
	if user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr) // not login
	}

	rdao := new(cache.WordDao)
	if rdao.IsLiked(c, wid, user.UID) {
		return sz.CodeResponse(sz.CodeWordDupLiked) // already liked
	}
	rdao.Like(c, wid, user.UID)

	resp = sz.Success()
	return
}

func (_ *WordLikeService) Unlike(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	wid := c.Param("id")
	if word, err := dao.QueryByID(wid); err != nil {
		zap.S().Warnf("query word error: %e", err)
		return sz.Error()
	} else if word == nil {
		return sz.Failure() // word not exists
	} else if word.IsDeleted() {
		// word have been deleted
		return sz.CodeResponse(sz.CodeWordDeleted)
	}
	user := model.CurrentUser(c)
	if user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr) // not login
	}

	rdao := new(cache.WordDao)
	if !rdao.IsLiked(c, wid, user.UID) {
		return sz.CodeResponse(sz.CodeWordNotLiked) // already liked
	}
	rdao.Unlike(c, wid, user.UID)

	resp = sz.Success()
	return
}

func (_ *WordLikeService) Trend(c *gin.Context) (resp sz.Response) {
	dao := new(database.WordDao)
	rdao := new(cache.WordDao)
	trend := rdao.Trend(c) // string slice
	words := make([]*model.Word, 0, len(trend))
	for _, v := range trend {
		word, err := dao.QueryByID(v)
		if err != nil || word == nil {
			continue
		}
		word.Likes = rdao.Likes(c, v)
		words = append(words, word)
	}
	resp = sz.Success()
	resp.Data = sz.NewTrend(words)
	return
}
