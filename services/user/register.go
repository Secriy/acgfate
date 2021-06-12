package user

import (
	"fmt"
	"time"

	"acgfate/log"
	"acgfate/model"
	"acgfate/model/user"
	sz "acgfate/serializer"
	suser "acgfate/serializer/user"
	_ "github.com/gin-gonic/gin"
)

type RegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Email    string `json:"email" binding:"required,email"`
}

// Register 用户注册服务
func (service *RegisterService) Register() sz.Response {
	var baseInfo user.BaseInfo
	// 判断用户名是否被占用
	if exist("username", service.Username) {
		return sz.ErrResponse(sz.RegNameExist)
	}
	// 判断邮箱是否被占用
	if exist("email", service.Email) {
		return sz.ErrResponse(sz.EmailExist)
	}
	// 加密密码
	if err := baseInfo.SetPassword(service.Password); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.PasswdEncryptErr), err)
		return sz.ErrResponse(sz.PasswdEncryptErr)
	}
	// 创建用户信息记录
	infoSQL := "INSERT INTO user_base_info (username, password, nickname, email, join_time) VALUES (?,?,?,?,?)"
	rows, err := model.DB.Exec(infoSQL, service.Username, baseInfo.Password, service.Nickname, service.Email, time.Now())
	if err != nil {
		log.Logger.Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}
	uid, _ := rows.LastInsertId()
	baseInfo, err = user.GetUserInfoByID(uid)
	if err != nil {
		log.Logger.Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}

	return sz.BuildResponse(
		sz.Success,
		suser.BuildBaseInfoResponse(&baseInfo),
		sz.Msg(sz.Success),
	)
}

// 判断是否存在字段
func exist(column string, arg string) bool {
	var count int
	row := model.DB.QueryRow(fmt.Sprintf(
		"SELECT 1 FROM user_base_info WHERE %s = ? LIMIT 1", column), arg)
	if _ = row.Scan(&count); count > 0 {
		msg := sz.Msg(sz.RegNameExist)
		log.Logger.Infof("%s: %s", msg, arg)
		return true
	}
	return false
}
