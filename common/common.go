package common

import (
	"Go_forum/models"
	"errors"
	"github.com/gin-gonic/gin"
)

func VRI(register *models.Register) bool {
	if len(register.Username) == 0 || len(register.Password) == 0 || len(register.Re_Password) == 0 || register.Password != register.Re_Password {
		return false
	}
	return true
}

func GetCurrentUserID(c *gin.Context) (user_id int64, err error) {
	uid, ok := c.Get("user_id")
	if !ok {
		err = errors.New("无用户id")
		return
	}
	user_id, ok = uid.(int64)
	if !ok {
		err = errors.New("无用户id.")
		return
	}
	return
}
