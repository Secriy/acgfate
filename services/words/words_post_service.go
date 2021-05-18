package words

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	Content string `json:"content" binding:"required,max=1000"`
}

func (service *PostService) Post(c *gin.Context) sz.Response {
	words := model.Words{
		Publisher: c.GetUint64("UID"),
		Content:   service.Content,
	}
	// 新增文字
	if err := model.DB.Create(&words).Error; err != nil {
		return sz.ErrorResponse(sz.WordsPostErr, sz.GetResMsg(sz.WordsPostErr))
	}
	// 获取发布者昵称
	user, err := model.GetUser(words.Publisher)
	if err != nil {
		return sz.ErrorResponse(sz.WordsPostErr, sz.GetResMsg(sz.WordsPostErr))
	}

	return sz.BuildResponse(sz.Success, sz.BuildWordsResponse(&words, user.Nickname), "发表成功")
}
