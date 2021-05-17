package words

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type GetService struct{}

func (service *GetService) Get(c *gin.Context) serializer.Response {
	var words model.Words
	if err := model.DB.First(&words, c.Param("wid")).Error; err != nil {
		return serializer.Error(utils.Error, "获取文字错误")
	}
	// 获取发布者昵称
	user, err := model.GetUser(words.Publisher)
	if err != nil {
		return serializer.Error(utils.WordsPostErr, utils.GetResMsg(utils.WordsPostErr))
	}

	return serializer.BuildResponse(utils.Success, serializer.BuildWordsResponse(&words, user.Nickname), "成功")
}
