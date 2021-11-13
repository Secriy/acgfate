package serializer

import (
	"strconv"

	"acgfate/model"
)

// Category 用户账号信息
type Category struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

// NewCategory 构建分区信息响应
func NewCategory(cat *model.Category) Category {
	return Category{
		CategoryID:   strconv.FormatInt(cat.CategoryID, 10),
		CategoryName: cat.CategoryName,
		Description:  cat.Description,
	}
}

// NewMultiCategory 构建多个分区信息响应
func NewMultiCategory(cat []*model.Category) []Category {
	ret := make([]Category, 0, len(cat))
	for _, v := range cat {
		ret = append(ret, Category{
			CategoryID:   strconv.FormatInt(v.CategoryID, 10),
			CategoryName: v.CategoryName,
			Description:  v.Description,
		})
	}
	return ret
}
