package words

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	Content string `json:"content" form:"content" binding:"required,max=1000"`
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

	return serializer.BuildResponse(utils.Success, serializer.BuildWordsResponse(&words), "发表成功")
}
