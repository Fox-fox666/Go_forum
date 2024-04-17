package controllers

import (
	"Go_forum/common"
	"Go_forum/logic"
	"Go_forum/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	user_id, err := common.GetCurrentUserID(c)
	if err != nil {
		ResErr(c, "用户未登录", nil)
		return
	}
	vd := &models.VoteData{UserID: user_id}
	err = c.ShouldBindJSON(vd)
	if err != nil {
		zap.L().Error("PostVote参数有误" + err.Error())

		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResErr(c, "", common.RemoveTopStruct(errs.Translate(common.Trans)))
			return
		}

		ResErr(c, err.Error(), nil)
		return
	}

	err = logic.PostVote(vd)
	if err != nil {
		zap.L().Error("logic.PostVote有误" + err.Error())
		ResErr(c, "server error", nil)
		return
	}

	ResSuccess(c, "", nil)
}
