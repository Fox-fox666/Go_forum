package reDis

import "fmt"

const (
	Prefix           = "Go_forum"
	KeyPostTime      = "post:time" //记录发帖时间
	KeyPostSocreZset = "post:score"
	KeyPostVotedZset = "post:voted:"
)

func getKey(Key string) string {
	return fmt.Sprintf("%s:%s", Prefix, Key)
}

//Zrange Go_forum:post:time 0 -1
