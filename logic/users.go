package logic

import (
	"Go_forum/common"
	mySql "Go_forum/dao/mysql"
	"Go_forum/models"
	"Go_forum/pkg"
	"database/sql"
	"errors"
)

func Register(register *models.Register) error {
	//判断用户名是否已经存在
	exist, err := mySql.CheckUserExist(register.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}

	//生成UID
	Userid := pkg.GenID()

	user := models.User{
		User_id:  Userid,
		Username: register.Username,
		Password: common.EncryptPassword(register.Password),
	}

	//保存进数据库
	err = mySql.InsertUser(&user)

	return err
}

func Login(login *models.Login) (token string, err error) {
	user := &models.User{
		User_id:  0,
		Username: login.Username,
		Password: "",
	}
	err = mySql.QueryUserByUsername(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("用户名或密码错误")
		}
		return "", err
	}

	if common.EncryptPassword(login.Password) != user.Password {
		return "", errors.New("用户名或密码错误")
	}

	token, err = pkg.GenToken(user.Username, user.User_id)
	if err != nil {
		return "", errors.New("系统登录错误")
	}
	return token, err
}
