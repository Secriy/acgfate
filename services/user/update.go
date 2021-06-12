package user

import (
	"time"

	"acgfate/log"
	"acgfate/model"
	"acgfate/model/user"
	sz "acgfate/serializer"
	suser "acgfate/serializer/user"
	"github.com/gin-gonic/gin"
)

type UpdateService struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=15"`
	Email    string `json:"email" binding:"omitempty,email"`
	Sign     string `json:"sign" binding:"omitempty,min=1,max=70"`
	Gender   string `json:"gender" binding:"omitempty"` // 只能为 男、女、其他、保密
	Birthday string `json:"birthday" binding:"omitempty"`
	Province string `json:"province" binding:"omitempty"`
	City     string `json:"city" binding:"omitempty"`
}

// Update 用户信息更新服务
func (service *UpdateService) Update(c *gin.Context) sz.Response {
	// 绑定用户模型
	var baseInfo user.BaseInfo
	if err := baseInfo.GetBaseInfo(c); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}
	// 判断字段是否为空
	if service.Nickname != "" {
		baseInfo.Nickname = service.Nickname
	}
	if service.Gender != "" {

		baseInfo.Gender = genderID(service.Gender)
	}
	if service.Sign != "" {
		baseInfo.Sign = service.Sign
	}
	if service.Birthday != "" {
		var err error
		baseInfo.Birthday, err = time.Parse("2006-01-02", service.Birthday)
		if err != nil {
			return sz.MsgResponse(sz.FormatErr, "生日格式不正确")
		}
	}
	if service.Province != "" {
		baseInfo.Province = service.Province
	}
	if service.City != "" {
		baseInfo.City = service.City
	}
	// 判断是否修改邮箱
	if service.Email != "" && baseInfo.Email != service.Email {
		baseInfo.Email = service.Email
		baseInfo.EmailVerified = false
	}
	// 更新数据
	sqlStr := "UPDATE user_base_info SET nickname=?, email=?, mail_verified=?, sign=?, gender=?, province=?, city=?, " +
		"birthday=? where uid=?"
	if _, err := model.DB.Exec(sqlStr, baseInfo.Nickname, baseInfo.Email, baseInfo.EmailVerified, baseInfo.Sign,
		baseInfo.Gender, baseInfo.Province, baseInfo.City,
		baseInfo.Birthday, baseInfo.UID); err != nil {
		log.Logger.Errorf("更新用户信息失败: %s", err)
		return sz.ErrResponse(sz.UpdateDBErr)
	}

	log.Logger.Infof("更新用户信息成功: %d", baseInfo.UID)

	return sz.BuildResponse(
		sz.Success,
		suser.BuildBaseInfoResponse(&baseInfo),
		sz.Msg(sz.Success),
	)
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
