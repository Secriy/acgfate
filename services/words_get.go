package services

import (
	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
)

type GetService struct{}

// Get 获取Words
func (service *GetService) Get(wid interface{}) sz.Response {
	var words model.Words
	if err := words.WidGet(wid); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.MsgResponse(sz.Error, "获取文字失败")
	}
	info, err := model.GetUserInfoByID(words.UID)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}

	return sz.BuildWordsResponse(&words, info.Nickname)
}
