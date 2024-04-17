package mySql

import (
	"Go_forum/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

func GetAllCommunityList() (communityList []*models.Community, err error) {
	sqlstr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlstr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no community in db")
			err = nil
		}
		zap.L().Error(err.Error())
	}
	return
}

func GetCommunityDetailById(communityid int64) (*models.CommunityDetail, error) {
	x := models.CommunityDetail{}
	sqlstr := "select community_id,community_name,introduction,create_time from community where community_id=?"
	err := db.Get(&x, sqlstr, communityid)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no community(id) in db")
		}
		return nil, err
	}
	return &x, nil
}

func GetCommunityById(cid int64) (c *models.Community, err error) {
	c = new(models.Community)
	sqlstr := "select community_id,community_name from community where community_id=?"
	err = db.Get(c, sqlstr, cid)
	if err == sql.ErrNoRows {
		err = errors.New("无该id社区")
		return nil, err
	}
	return
}
