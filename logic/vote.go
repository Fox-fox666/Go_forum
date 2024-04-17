package logic

import (
	reDis "Go_forum/dao/redis"
	"Go_forum/models"
	"errors"
	"go.uber.org/zap"
	"math"
	"time"
)

/*投票分析

1：没投票的投赞成，投反对的改投赞成
0：取消投票
-1：没投票的投反对，投赞成的改投反对

限制：
每个帖子发表之后只有一周之内才能投票，超过一周的不允许投票
    到期后，赞成票和反对票存到mysql
    到期后，redis中删除该帖子对应的投票表（PostVotedZset）
*/

const (
	oneWeekTimeSeconds = 7 * 3600 * 24
	scorePerVote       = 412 //每一票的分数
)

func PostVote(data *models.VoteData) error {
	//判断帖子是否还可以投票
	postCreateTime := reDis.GetPostCreateTime(data.PostId)
	if float64(time.Now().Unix())-postCreateTime > oneWeekTimeSeconds {
		err := reDis.DelZset(reDis.KeyPostVotedZset)
		if err != nil {
			zap.L().Error("DelZset error")
			return err
		}
		return errors.New("帖子过期，无法投票")
	}
	//更新帖子分数
	//先查看之前的投票数据
	votedYN := reDis.GetPostVoted(data.UserID, data.PostId)
	diff := math.Abs(votedYN - float64(data.YesNo))
	postScore := reDis.GetPostScore(data.PostId)

	var updatedScore float64
	if data.YesNo != 0 {
		updatedScore = postScore + diff*float64(data.YesNo)*scorePerVote
	} else {
		updatedScore = postScore + votedYN*-1*scorePerVote
	}

	err := reDis.UpdatePostScore(data.PostId, updatedScore)
	if err != nil {
		return err
	}

	//记录用户为帖子投票的数据
	err = reDis.UpdatePostVoted(data.UserID, data.PostId, data.YesNo)
	if err != nil {
		return err
	}

	return nil
}
