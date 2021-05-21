package user

import (
	"fmt"

	"acgfate/model"
	sz "acgfate/serializer"
)

type RegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Mail     string `json:"mail" binding:"required,email"`
}

// Register 用户注册服务
func (service RegisterService) Register() sz.Response {
	var userInfo = model.UserInfo{
		Username: service.Username,
		Password: service.Password,
		Nickname: service.Nickname,
		Mail:     service.Mail,
	}
	// 判断用户名是否已经存在
	sqlStr := "SELECT * FROM user_info WHERE username = ?"
	var user model.UserInfo
	err := model.DB.Get(&user, sqlStr, service.Username)
	if err == nil && !user.IsDeleted {
		return sz.ErrResponse(sz.Failure, "用户名已被他人使用")
	}
	// 加密密码
	if err = userInfo.SetPassword(service.Password); err != nil {
		return sz.ErrResponse(sz.CodeEncryptError, "密码加密失败")
	}
	// 创建用户信息记录
	sqlInfoStr := "INSERT INTO user_info (username, password, nickname, mail) VALUES (?,?,?,?)"
	rows, err := model.DB.Exec(sqlInfoStr, userInfo.Username, userInfo.Password, userInfo.Nickname, userInfo.Mail)
	if err != nil {
		return sz.ErrResponse(sz.DatabaseErr, "创建用户失败")
	}
	// 创建用户点数记录
	sqlPointStr := "INSERT INTO user_point (exp, coins) VALUES (DEFAULT, DEFAULT)"
	_, err = model.DB.Exec(sqlPointStr)
	if err != nil {
		return sz.ErrResponse(sz.DatabaseErr, "创建用户失败")
	}
	// 获取刚刚创建用户的ID
	uid, err := rows.LastInsertId()
	if err != nil {
		return sz.ErrResponse(sz.DatabaseErr, "获取用户ID错误")
	}
	fmt.Println(uid)
	// 构建模型
	info, err := model.GetUserInfo(uid)
	if err != nil {
		return sz.ErrResponse(sz.Failure, "创建模型失败")
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildUserInfoResponse(&info),
		sz.GetResMsg(sz.Success),
	)
}
