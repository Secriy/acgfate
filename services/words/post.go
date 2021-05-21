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
	uid := c.GetUint64("UID")
	// 新增文字
	sql := "INSERT INTO words (publisher, content) VALUES (?,?)"
	rows, err := model.DB.Exec(sql, uid, service.Content)
	if err != nil {
		return sz.ErrResponse(sz.DatabaseErr, "创建文字失败")
	}
	wid, err := rows.LastInsertId()
	if err != nil {
		return sz.ErrResponse(sz.DatabaseErr, "创建文字失败")
	}
	// 获取发布者昵称
	user, err := model.GetUserInfo(uid)
	if err != nil {
		return sz.ErrResponse(sz.WordsPostErr, "获取用户失败")
	}
	// 获取Words模型
	words, err := model.GetWordsByWID(wid)
	if err != nil {
		return sz.ErrResponse(sz.WordsPostErr, "获取文字失败")
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildWordsResponse(&words, user.Nickname),
		"发表成功",
	)
}
