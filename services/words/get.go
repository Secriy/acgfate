package words

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type GetService struct{}

func (service *GetService) Get(c *gin.Context) sz.Response {
	words, err := model.GetWordsByWID(c.Param("wid"))
	if err != nil {
		return sz.ErrResponse(sz.Error, "获取文字错误")
	}
	// 获取发布者昵称
	user, err := model.GetUserInfo(words.Publisher)
	if err != nil {
		return sz.ErrResponse(sz.Error, "获取发布者信息错误")
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildWordsResponse(&words, user.Nickname),
		"成功",
	)
}
