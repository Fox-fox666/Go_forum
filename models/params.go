package models

//定义请求的参数结构体

type Register struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Re_Password string `json:"re_password" binding:"required,eqfield=Password"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
