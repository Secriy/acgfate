package serializer

import (
	"acgfate/model"
)

// CateResponse 分区响应结构
type CateResponse struct {
	CateID   int64  `json:"cid"`
	CateName string `json:"cname"`
	Desc     string `json:"desc"`
}

// BuildCateResponse 构建分区信息响应
func BuildCateResponse(cate *model.Category) interface{} {
	return CateResponse{
		CateID:   cate.CateID,
		CateName: cate.CateName,
		Desc:     cate.Desc,
	}
}

// BuildCateMultiResponse 构建多个分区信息响应
func BuildCateMultiResponse(cates []*model.Category) interface{} {
	var rets []interface{}
	for _, v := range cates {
		cate := BuildCateResponse(v)
		rets = append(rets, cate)
	}
	return rets
}
