package routers

import (
	"Go_forum/controllers"
	"Go_forum/logger"
	"Go_forum/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup() *gin.Engine {
	if viper.GetString("app.mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	//中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	v1 := r.Group("/api/v1")
	//路由信息
	v1.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//注册
	v1.POST("/register", controllers.RegisterHandler)

	//登录
	v1.POST("/login", controllers.Login)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetail)
		v1.GET("/post", controllers.GetPostList)
		//按照分数或者时间排序获取帖子列表
		v1.GET("/post2",controllers.GetPostListBySome)

		v1.POST("/vote", controllers.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
