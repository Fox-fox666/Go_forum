package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	//viper.SetConfigFile("config.yaml")
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 指定配置文件类型（专用于远程获取配置信息时指定配置类型）
	viper.AddConfigPath("./conf") //指定查找配置文件的路径

	err := viper.ReadInConfig() // 查找并读取配置文件

	if err != nil { // 处理读取配置文件的错误
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("viper.ConfigFileNotFoundError")
			return err
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("Fatal error config file:", err)
			return err
		}
	}
	// 配置文件找到并成功解析
	//实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//配置文件发生变化之后会调用的回调函数
		fmt.Println("Config file changed", in.Name)
	})
	return nil
}
