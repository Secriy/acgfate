package serializer

import (
	"acgfate/model"
)

type Trend struct {
	Rank int `json:"rank"`
	Word
}

// NewTrend 构建排行响应
func NewTrend(words []*model.Word) (ret []Trend) {
	ret = make([]Trend, 0, len(words))
	for k, v := range words {
		ret = append(ret, Trend{
			Rank: k + 1,
			Word: NewWord(v),
		})
	}
	return
}
