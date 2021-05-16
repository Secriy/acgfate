package v1

import (
	"acgfate/services/user"
	"github.com/gin-gonic/gin"
)

// UserRegister is the API of user register
func UserRegister(c *gin.Context) {
	var form user.RegisterService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(40001, err))
	}

}

// UserLogin is the API of user login
func UserLogin(c *gin.Context) {
	var form user.LoginService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Login(c)
		c.JSON(200, res)
	} else {

		c.JSON(200, ErrorResponse(40001, err))
	}
}
