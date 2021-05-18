package words

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type GetService struct{}

func (service *GetService) Get(c *gin.Context) sz.Response {
	var words model.Words
	if err := model.DB.First(&words, c.Param("wid")).Error; err != nil {
		return sz.ErrorResonse(sz.Error, "获取文字错误")
	}
	// 获取发布者昵称
	user, err := model.GetUser(words.Publisher)
	if err != nil {
		return sz.ErrorResonse(sz.WordsPostErr, sz.GetResMsg(sz.WordsPostErr))
	}

	return sz.BuildResponse(sz.Success, sz.BuildWordsResponse(&words, user.Nickname), "成功")
}
