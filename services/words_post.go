package services

import (
	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	Title    string `json:"title" binding:"omitempty,min=1,max=20"`
	Content  string `json:"content" binding:"required,min=1,max=10000"`
	Category uint8  `json:"category" binding:"required"`
	Tags     string `json:"tags" binding:"omitempty,min=1,max=90"`
}

func (service *PostService) Post(c *gin.Context) sz.Response {
	acc := model.CurrentUser(c)
	// 创建用户账号记录
	accSQL := "INSERT INTO words (uid, title, content, category, tags) VALUES (?,?,?,?,?)"
	_, err := model.DB.Exec(accSQL, acc.UID, service.Title, service.Content, service.Category, service.Tags)
	if err != nil {
		log.Logger.Errorf("创建Words失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建Words失败")
	}
	return sz.SuccessResponse()
}
