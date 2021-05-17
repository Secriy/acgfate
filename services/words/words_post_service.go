package words

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	Content string `json:"content" binding:"required,max=1000"`
}

func (service *PostService) Post(c *gin.Context) serializer.Response {
	words := model.Words{
		Publisher: c.GetUint64("UID"),
		Content:   service.Content,
	}
	// 新增文字
	if err := model.DB.Create(&words).Error; err != nil {
		return serializer.Error(utils.WordsPostErr, utils.GetResMsg(utils.WordsPostErr))
	}
	// 获取发布者昵称
	user, err := model.GetUser(words.Publisher)
	if err != nil {
		return serializer.Error(utils.WordsPostErr, utils.GetResMsg(utils.WordsPostErr))
	}

	return serializer.BuildResponse(utils.Success, serializer.BuildWordsResponse(&words, user.Nickname), "发表成功")
}
