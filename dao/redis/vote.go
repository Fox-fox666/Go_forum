package reDis

import (
	"Go_forum/models"
	"github.com/go-redis/redis"
	"strconv"
)

func PostVote(data *models.VoteData) {

}

func GetPostCreateTime(pid int64) float64 {
	pidStr := strconv.FormatInt(pid, 10)
	return rdb.ZScore(getKey(KeyPostTime), pidStr).Val()
}

// 得到uid用户给pid帖子投的什么票
func GetPostVoted(uid int64, pid int64) float64 {
	pidStr := strconv.FormatInt(pid, 10)
	uidStr := strconv.FormatInt(uid, 10)
	return rdb.ZScore(getKey(KeyPostVotedZset+pidStr), uidStr).Val()
}

func UpdatePostScore(pid int64, score float64) error {
	pidStr := strconv.FormatInt(pid, 10)
	return rdb.ZAdd(getKey(KeyPostSocreZset), redis.Z{
		Score:  score,
		Member: pidStr,
	}).Err()
}

func GetPostScore(pid int64) float64 {
	pidStr := strconv.FormatInt(pid, 10)
	return rdb.ZScore(getKey(KeyPostSocreZset), pidStr).Val()
}

func UpdatePostVoted(uid int64, pid int64, yn int8) error {
	pidStr := strconv.FormatInt(pid, 10)
	uidStr := strconv.FormatInt(uid, 10)
	if yn == 0 {
		return rdb.ZRem(getKey(KeyPostVotedZset+pidStr), uidStr).Err()
	}
	return rdb.ZAdd(getKey(KeyPostVotedZset+pidStr), redis.Z{
		Score:  float64(yn),
		Member: uidStr,
	}).Err()
}

func DelZset(key string) error {
	return rdb.Del(getKey(key)).Err()
}

func UpdatePostTime(pid int64, updatedTime float64) error {
	pidStr := strconv.FormatInt(pid, 10)
	return rdb.ZAdd(getKey(KeyPostTime), redis.Z{
		Score:  float64(updatedTime),
		Member: pidStr,
	}).Err()
}

func GetPostIdListBy(order string, page int64, size int64) (Idslice []int64, err error) {
	start := size * (page - 1)
	end := start + (page - 1)
	switch order {
	case "score":
		revRange, err := rdb.ZRevRange(getKey(KeyPostSocreZset), start, end).Result()
		if err != nil {
			return nil, err
		}
		Idslice = make([]int64, len(revRange))
		for i, _ := range revRange {
			Idslice[i], _ = strconv.ParseInt(revRange[i], 10, 64)
		}
		break
	}
	return
}

func GetVoteDataByIds(ids []int64) (votedata []int64, err error) {
	//使用pipeline一次发送多条命令减少RTT
	pipe := rdb.Pipeline()
	for _, id := range ids {
		pipe.ZCount(getKey(KeyPostVotedZset+strconv.FormatInt(id, 10)), "1", "1")
	}
	exec, err := pipe.Exec()
	if err != nil {
		return nil, err
	}
	votedata = make([]int64, len(exec))
	for i, cmder := range exec {
		votedata[i] = cmder.(*redis.IntCmd).Val()
	}
	return
}
