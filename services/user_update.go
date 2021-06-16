package services

import (
	"time"

	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type UpdateInfoService struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=15"`
	Sign     string `json:"sign" binding:"omitempty,min=1,max=70"`
	Gender   string `json:"gender" binding:"omitempty"` // 只能为 男、女、其他、保密
	Birthday string `json:"birthday" binding:"omitempty"`
	Province string `json:"province" binding:"omitempty"`
	City     string `json:"city" binding:"omitempty"`
}

// Update 用户信息更新服务
func (service *UpdateInfoService) Update(c *gin.Context) sz.Response {
	// 绑定用户模型
	var basicInfo model.BasicInfo
	username, err := basicInfo.CurrentBasicInfo(c)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}
	// 判断字段是否为空
	basicInfo.Nickname = getNotNone(basicInfo.Nickname, service.Nickname) // 昵称
	basicInfo.Sign = getNotNone(basicInfo.Sign, service.Sign)             // 签名
	if service.Gender != "" {                                             // 性别
		basicInfo.Gender = genderID(service.Gender)
	}
	if service.Birthday != "" { // 生日
		basicInfo.Birthday, err = time.Parse("2006-01-02", service.Birthday)
		if err != nil {
			return sz.MsgResponse(sz.FormatErr, "生日格式不正确")
		}
	}
	basicInfo.Province = getNotNone(basicInfo.Province, service.Province) // 省份
	basicInfo.City = getNotNone(basicInfo.City, service.City)             // 城市
	// 更新数据
	sqlStr := "UPDATE user_basic_info SET nickname=?, sign=?, gender=?, province=?, city=?, birthday=? WHERE uid=?"
	if _, err = model.DB.Exec(sqlStr, basicInfo.Nickname, basicInfo.Sign, basicInfo.Gender, basicInfo.Province, basicInfo.City,
		basicInfo.Birthday, basicInfo.UID); err != nil {
		log.Logger.Errorf("更新用户信息失败: %s", err)
		return sz.ErrResponse(sz.UpdateDBErr)
	}

	log.Logger.Infof("更新用户信息成功: %d", basicInfo.UID)

	return sz.BuildBaseInfoResponse(&basicInfo, username)
}

func getNotNone(old, new string) string {
	if new != "" {
		return new
	}
	return old
}

func genderID(sex string) uint8 {
	switch sex {
	case "男":
		return 1
	case "女":
		return 2
	case "其他":
		return 3
	default:
		return 0
	}
}
