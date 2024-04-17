package mySql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	fmt.Println(dsn)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect MySQL DB failed, err", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("max_open_connection"))
	db.SetMaxIdleConns(viper.GetInt("max_idle_connection"))
	return
}

func Close() {
	err := db.Close()
	if err != nil {
		fmt.Println("ddddd")
	}
}
