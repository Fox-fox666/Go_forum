package models

type VoteData struct {
	UserID int64 `json:"user_id,string"`                    //谁点赞
	PostId int64 `json:"post_id,string" binding:"required"` //帖子ID
	YesNo  int8  `json:"yes_no" binding:"oneof=-1 0 1"`     //点赞（1）还是👎（-1）,还是取消点赞（0）
}
