package controllers

import (
	"Go_forum/common"
	"Go_forum/logic"
	"Go_forum/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func CreatePostHandler(c *gin.Context) {
	//获取参数以及参数校验
	post := new(models.Post)
	post.CreateTime = time.Now()
	err := c.ShouldBind(post)
	if err != nil {
		zap.L().Error("Post参数有误")
		ResErr(c, "Post参数有误", nil)
		return
	}

	//创建帖子
	user_id, err := common.GetCurrentUserID(c)
	if err != nil {
		ResErr(c, "用户未登录", nil)
		return
	}
	post.AuthorID = user_id
	err = logic.CreatePost(post)
	if err != nil {
		zap.L().Error("CreatePost error:" + err.Error())
		ResErr(c, "server error", nil)
		return
	}
	//返回响应
	ResSuccess(c, "创建成功", nil)
}

// 获取帖子详情
func GetPostDetail(c *gin.Context) {
	pidstr := c.Param("id")
	pid, err := strconv.ParseInt(pidstr, 10, 64)
	if err != nil {
		ResErr(c, "postid参数有误", nil)
		return
	}

	postdetail, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("GetPostDetailById failed:" + err.Error())
		ResErr(c, "server error", nil)
		return
	}
	ResSuccess(c, "", postdetail)
}

func GetPostList(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	postList, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList err" + err.Error())
		ResErr(c, "server error", nil)
		return
	}
	ResSuccess(c, "", postList)
}

//redis拿id列表
//根据id去mysql拿帖子详情

func GetPostListBySome(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	order := c.Query("order")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	postList, err := logic.GetPostIdListBy(order, page, size)
	if err != nil {
		zap.L().Error("logic.GetPostIdListBy err" + err.Error())
		ResErr(c, "server error", nil)
		return
	}
	ResSuccess(c, "", postList)
}
