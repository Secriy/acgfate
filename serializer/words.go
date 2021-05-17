package serializer

import (
	"time"

	"acgfate/model"
)

type WordsResponse struct {
	WID       uint64    `json:"wid"`
	Publisher string    `json:"publisher"`
	Content   string    `json:"content"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BuildWordsResponse 文字信息返回构建
func BuildWordsResponse(words *model.Words, nickname string) WordsResponse {
	return WordsResponse{
		WID:       words.WID,
		Publisher: nickname,
		Content:   words.Content,
		CreateAt:  words.CreatedAt,
		UpdatedAt: words.UpdatedAt,
	}
}
