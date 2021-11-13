package serializer

import "acgfate/model"

// Category 用户账号信息
type Category struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

// NewCategory 构建分区信息响应
func NewCategory(cate *model.Category) Category {
	return Category{
		CategoryID:   cate.CategoryID,
		CategoryName: cate.CategoryName,
		Description:  cate.Description,
	}
}

// NewMultiCategory 构建多个分区信息响应
func NewMultiCategory(cate []*model.Category) []Category {
	ret := make([]Category, len(cate))
	for k, v := range cate {
		ret[k] = Category{
			CategoryID:   v.CategoryID,
			CategoryName: v.CategoryName,
			Description:  v.Description,
		}
	}
	return ret
}
