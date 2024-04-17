package mySql

import (
	"Go_forum/models"
	"database/sql"
	"errors"
)

func CreatePost(post *models.Post) error {
	sqlStr := `insert into post 
		(post_id,title,content,author_id,community_id,status)
		values (?,?,?,?,?,?) 
		`
	_, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID, post.Status)
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(pid int64) (p *models.Post, err error) {
	sqlStr := `select
              post_id,title,content,author_id,community_id,status,create_time
			 from post
             WHERE post_id=?
             `
	p = new(models.Post)
	err = db.Get(p, sqlStr, pid)
	if err == sql.ErrNoRows {
		err = errors.New("无该id贴")
	}
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
              post_id,title,content,author_id,community_id,status,create_time
			 from post
			 ORDER BY create_time
			 DESC 
             limit ?,?
             `
	posts = make([]*models.Post, 0, size)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
