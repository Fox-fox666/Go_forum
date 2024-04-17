package controllers

import (
	"Go_forum/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	//查询所有的社区（community_id,community_name）以列表形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed")
		ResErr(c, "server error", nil)
		return
	}
	ResSuccess(c, "", communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	communityId, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResErr(c, "参数有误", nil)
		return
	}
	communityDetail, err := logic.GetCommunityDetail(communityId)
	if err != nil {
		ResErr(c, "未找到对应社区", nil)
		return
	}
	ResSuccess(c, "", communityDetail)
}
