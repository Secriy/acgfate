package serializer

import (
	"time"

	"acgfate/model"
)

type WordsResponse struct {
	WID      uint64    `json:"wid"`
	Author   string    `json:"author"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Views    uint      `json:"views"`
	Likes    uint      `json:"likes"`
	Comments uint      `json:"comments"`
	Category string    `json:"category"`
	Tags     string    `json:"tags"`
	PostTime time.Time `json:"post_time"`
	EditTime time.Time `json:"edit_time"`
}

// BuildWordsResponse 构建用户信息响应
func BuildWordsResponse(w *model.Words, nickname string) Response {
	return BuildResponse(Success, WordsResponse{
		WID:      w.WID,
		Author:   nickname,
		Title:    w.Title,
		Content:  w.Content,
		Views:    w.Views,
		Likes:    w.Likes,
		Comments: w.Comments,
		Category: tagDesc(w.Category),
		Tags:     w.Tags,
		PostTime: w.CreatedAt,
		EditTime: w.UpdatedAt,
	}, Msg(Success))
}
