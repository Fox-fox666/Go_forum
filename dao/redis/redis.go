package reDis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明全局的rdb变量
var rdb *redis.Client

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"), //连接池大小
	})

	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	err := rdb.Close()
	if err != nil {
		fmt.Println("rrrr")
	}
}
