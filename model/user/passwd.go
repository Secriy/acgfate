package user

import "golang.org/x/crypto/bcrypt"

// SetPassword 设置密码
func (info *BaseInfo) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)
	if err != nil {
		return err
	}
	info.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (info *BaseInfo) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	return err == nil
}
