package serializer

import (
	"acgfate/database"
	"acgfate/model"
	"acgfate/util"
)

// Word 文字模型
type Word struct {
	Wid      int64  `json:"wid"`
	Author   string `json:"author"`
	Category string `json:"category"`
	// Status    string    `json:"status"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedTime string `json:"created_time"`
	EditedTime  string `json:"edited_time"`
}

func NewWord(word *model.Word) Word {
	return Word{
		Wid:         word.Wid,
		Author:      database.NickName(word.Aid),
		Category:    database.CatName(word.CatID),
		Title:       word.Title,
		Content:     word.Content,
		CreatedTime: util.TimeFormat(word.CreatedAt),
		EditedTime:  util.TimeFormat(word.UpdatedAt),
	}
}

// NewMultiWord 构建多个文字信息响应
func NewMultiWord(words []*model.Word) []Word {
	ret := make([]Word, 0, len(words))
	for _, v := range words {
		ret = append(ret, Word{
			Wid:         v.Wid,
			Author:      database.NickName(v.Aid),
			Category:    database.CatName(v.CatID),
			Title:       v.Title,
			Content:     v.Content,
			CreatedTime: util.TimeFormat(v.CreatedAt),
			EditedTime:  util.TimeFormat(v.UpdatedAt),
		})
	}
	return ret
}
