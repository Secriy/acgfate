package serializer

// 分区对照表
var _wordsCategory = map[uint8]string{
	0: "默认",
	1: "日常",
	2: "字句",
	3: "讨论",
}

// tagDesc 返回分区名
func tagDesc(tagID uint8) string {
	if v, ok := _wordsCategory[tagID]; ok {
		return v
	}
	return _wordsCategory[0]
}
