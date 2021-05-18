package serializer

import (
	"acgfate/model"
)

type WordsResponse struct {
	WID       uint64 `json:"wid"`
	Publisher string `json:"publisher"`
	Content   string `json:"content"`
	CreateAt  int64  `json:"create_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// BuildWordsResponse 文字信息返回构建
func BuildWordsResponse(words *model.Words, nickname string) WordsResponse {
	return WordsResponse{
		WID:       words.WID,
		Publisher: nickname,
		Content:   words.Content,
		CreateAt:  words.CreatedAt.Unix(),
		UpdatedAt: words.UpdatedAt.Unix(),
	}
}
