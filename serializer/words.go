package serializer

import (
	"time"

	"acgfate/model"
)

type WordsResponse struct {
	Publisher uint64    `json:"publisher"`
	Content   string    `json:"content"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BuildWordsResponse 文字信息返回构建
func BuildWordsResponse(words *model.Words) WordsResponse {
	return WordsResponse{
		Publisher: words.Publisher,
		Content:   words.Content,
		CreateAt:  words.CreatedAt,
		UpdatedAt: words.UpdatedAt,
	}
}
