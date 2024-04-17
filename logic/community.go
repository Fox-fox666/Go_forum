package logic

import (
	mySql "Go_forum/dao/mysql"
	"Go_forum/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mySql.GetAllCommunityList()
}

func GetCommunityDetail(communityid int64) (*models.CommunityDetail, error) {
	return mySql.GetCommunityDetailById(communityid)
}
