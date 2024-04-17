package logic

import (
	mySql "Go_forum/dao/mysql"
	reDis "Go_forum/dao/redis"
	"Go_forum/models"
	snowflake "Go_forum/pkg"
	"errors"
	"time"
)

func CreatePost(post *models.Post) error {
	post.ID = snowflake.GenID()
	err := mySql.CreatePost(post)
	if err != nil {
		return err
	}
	err = reDis.UpdatePostTime(post.ID, float64(time.Now().Unix()))
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(pid int64) (*models.Post, error) {
	return mySql.GetPostById(pid)
}

func GetPostDetailById(pid int64) (*models.PostDetail, error) {
	post, err := GetPostById(pid)
	if err != nil {
		return nil, errors.New("GetPostById failed" + err.Error())
	}

	AuthorName, err := mySql.GetUsernameByUserid(post.AuthorID)
	if err != nil {
		return nil, errors.New("GetUsernameByUserid failed" + err.Error())
	}

	c, err := mySql.GetCommunityById(post.CommunityID)
	if err != nil {
		return nil, errors.New("GetCommunityById failed" + err.Error())
	}

	return &models.PostDetail{
		AuthorName: AuthorName,
		Post:       post,
		Community:  c,
	}, nil
}

func GetPostList(page, size int64) ([]*models.PostDetail, error) {
	postList, err := mySql.GetPostList(page, size)
	if err != nil {
		return nil, errors.New("mySql.GetPostList failed" + err.Error())
	}

	PostDetailList := make([]*models.PostDetail, 0, len(postList))
	for _, post := range postList {
		AuthorName, err := mySql.GetUsernameByUserid(post.AuthorID)
		if err != nil {
			return nil, errors.New("GetUsernameByUserid failed" + err.Error())
		}

		c, err := mySql.GetCommunityById(post.CommunityID)
		if err != nil {
			return nil, errors.New("GetCommunityById failed" + err.Error())
		}

		PostDetailList = append(PostDetailList, &models.PostDetail{
			AuthorName: AuthorName,
			Post:       post,
			Community:  c,
		})
	}

	return PostDetailList, err
}

func GetPostIdListBy(order string, page int64, size int64) ([]*models.PostDetail, error) {
	PostIdList, err := reDis.GetPostIdListBy(order, page, size)
	if err != nil {
		return nil, err
	}
	postDetails := make([]*models.PostDetail, len(PostIdList))
	votedata, err := reDis.GetVoteDataByIds(PostIdList)
	if err != nil {
		return nil, err
	}

	for i, id := range PostIdList {
		postDetails[i], err = GetPostDetailById(id)
		postDetails[i].VoteNum=votedata[i]
		if err != nil {
			return nil, err
		}
	}
	return postDetails, nil
}
