package services

import (
	"fmt"

	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
)

// Exist 判断是否存在字段
func Exist(table, column, arg string) bool {
	var count int
	row := model.DB.QueryRow(fmt.Sprintf(
		"SELECT 1 FROM %s WHERE %s = ? LIMIT 1", table, column), arg)
	if _ = row.Scan(&count); count > 0 {
		msg := sz.Msg(sz.RegNameExist)
		log.Logger.Infof("%s: %s", msg, arg)
		return true
	}
	return false
}
