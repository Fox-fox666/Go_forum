package mySql

import (
	"Go_forum/models"
	"database/sql"
	"errors"
)

//把每一步数据库操作封装成函数
//等待logic业务层调用

func InsertUser(user *models.User) error {
	sql := "insert into user (user_id, username, password) VALUES (?,?,?)"
	_, err := db.Exec(sql, user.User_id, user.Username, user.Password)
	return err
}

func CheckUserExist(Username string) (bool, error) {
	sql := "select count(user_id) from user where username=?"
	var count int
	err := db.Get(&count, sql, Username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryUserByUsername(user *models.User) error {
	sql := "select user_id,username,password from user where username=?"
	err := db.Get(user, sql, user.Username)
	return err
}

func GetUsernameByUserid(uid int64) (username string, err error) {
	username = ""
	sqlStr := `select username from user where user_id=?`
	err = db.Get(&username, sqlStr, uid)
	if err == sql.ErrNoRows {
		return "", errors.New("无该id用户")
	}
	return
}
