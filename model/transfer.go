package model

const (
	Private = iota // 性别: 保密
	Male           // 性别: 男
	Female         // 性别: 女
	Other          // 性别: 其他
)

const (
	StatusNormal = iota // 状态: 正常
	StatusBanned        // 状态: 被封禁
)

// 状态对照表
var _statusFlags = map[uint8]string{
	StatusNormal: "账号正常",
	StatusBanned: "账号被封禁",
}

// 等级对照表
var _userLevel = []int{
	0,    // 0
	10,   // 1
	220,  // 2
	450,  // 3
	890,  // 4
	1500, // 5
	2900, // 6
	4700, // 7
	7800, // 8
	9600, // 9
}

// GetLevelByExp 根据经验值获取用户等级
func GetLevelByExp(exp int) (level int) {
	for rLevel, v := range _userLevel {
		if exp < v {
			level = rLevel - 1
			break
		}
	}
	return
}
