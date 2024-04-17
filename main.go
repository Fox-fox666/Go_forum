package main

import (
	"Go_forum/common"
	mySql "Go_forum/dao/mysql"
	reDis "Go_forum/dao/redis"
	"Go_forum/logger"
	snowflake "Go_forum/pkg"
	"Go_forum/routers"
	"Go_forum/settings"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Go web 开发通用脚手架

func init() {
	//1、加载配置
	err := settings.Init()
	if err != nil {
		fmt.Println("settings init failed:", err)
	}

	//2、初始化日志
	err = logger.Init(viper.GetString("app.mode"))
	if err != nil {
		fmt.Println("logger init failed:", err)
	}
	zap.L().Debug("logger init success...")

	//3、初始化数据库连接
	//    3.1、MySQL
	err = mySql.Init()
	if err != nil {
		fmt.Println("mysql init failed:", err)
	}

	//	  3.2、Redis
	err = reDis.Init()
	if err != nil {
		fmt.Println("redis init failed:", err)
	}

	//雪花id模块初始化
	err = snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.mechine_id"))
	if err != nil {
		fmt.Println("snowflake init failed:", err)
	}

	//初始化vaildator翻译器
	err = common.InitTrans("zh")
	if err != nil {
		fmt.Println("Trans init failed:", err)
	}
}

func Run() {
	//4、注册路由
	r := routers.Setup()
	//5、启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

func main() {
	defer func() {
		err := zap.L().Sync()
		if err != nil {
			fmt.Println("zzzz")
		}
	}()
	defer mySql.Close()
	defer reDis.Close()
	Run()
}
