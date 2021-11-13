package service

type UserUpdateService struct {
	Gender   string `json:"gender" binding:""`
	Sign     string `json:"sign" binding:""`
	Birthday string `json:"birthday"`
	Province string `json:"province"`
	City     string `json:"city"`
}

// TODO: user update
