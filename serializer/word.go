package serializer

import (
	"strconv"

	"acgfate/database"
	"acgfate/model"
	"acgfate/util"
)

// Word 文字模型
type Word struct {
	Wid         string `json:"wid"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedTime string `json:"created_time"`
	EditedTime  string `json:"edited_time"`
}

func NewWord(word *model.Word) (ret Word) {
	ret = Word{
		Wid:         strconv.FormatInt(word.Wid, 10),
		Author:      database.NickName(word.Aid),
		Category:    database.CatName(word.CatID),
		Title:       word.Title,
		Content:     word.Content,
		CreatedTime: util.TimeFormat(word.CreatedAt),
		EditedTime:  util.TimeFormat(word.UpdatedAt),
	}
	if word.Status == model.StatusWordDeleted {
		ret.Title = "已被删除"
		ret.Content = "已被删除"
	}
	return
}

// NewMultiWord 构建多个文字信息响应
func NewMultiWord(words []*model.Word) (ret []Word) {
	ret = make([]Word, 0, len(words))
	for _, v := range words {
		ret = append(ret, NewWord(v))
	}
	return
}
