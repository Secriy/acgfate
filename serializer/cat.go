package serializer

import "acgfate/model"

// Category 用户账号信息
type Category struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

// NewCategory 构建分区信息响应
func NewCategory(cat *model.Category) Category {
	return Category{
		CategoryID:   cat.CategoryID,
		CategoryName: cat.CategoryName,
		Description:  cat.Description,
	}
}

// NewMultiCategory 构建多个分区信息响应
func NewMultiCategory(cat []*model.Category) []Category {
	ret := make([]Category, len(cat))
	for k, v := range cat {
		ret[k] = Category{
			CategoryID:   v.CategoryID,
			CategoryName: v.CategoryName,
			Description:  v.Description,
		}
	}
	return ret
}
