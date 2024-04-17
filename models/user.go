package models

type User struct {
	User_id  int64  `db:"user_id,string"`
	Username string `db:"username"`
	Password string `db:"password"`
}
