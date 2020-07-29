package models

import (
	"time"
)

type Post struct {
	Id        		int64		`json:"id"`
	Title     		string		`xorm:"title" json:"title"`
	Content  		string		`xorm:"conte" json:"content"`
	CreatedAt 		time.Time	`xorm:"created_at" json:"created_at"`
	Creator  		string		`xorm:"creator" json:"creator"`
	LikesNumber		int			`xorm:"likes_number" json:"likes_number"`
	CommentsNumber	int			`xorm:"comments_number" json:"comments_number"`
	UserId			int64		`xorm:"user_id" json:"user_id"`
}

type UserLikePost struct {
	UserId			int64
	PostId			int64
	CreatedAt		time.Time
}

func (user *User) CreatePost(post Post) (err error) {

	if err = Engine.Ping(); err != nil {
		return
	}

	if err = Engine.Sync2(&Post{}); err != nil {
		return
	}

	_, err = Engine.Insert(&Post{
		Id:             0,
		Title:          post.Title,
		Content:        post.Content,
		CreatedAt:      time.Now(),
		Creator:        user.Name,
		LikesNumber:    0,
		CommentsNumber: 0,
		UserId:			user.Id,
	})
	return
}

func (user *User) PostLike(postId int64) (err error) {

	if err = Engine.Ping(); err != nil {
		return
	}

	if err = Engine.Sync2(&Post{}); err != nil {
		return
	}

	_, err = Engine.Insert(&UserLikePost{
		UserId: 	user.Id,
		PostId: 	postId,
		CreatedAt:	time.Now(),
	})

	if err != nil {
		return
	}

	err, post := PostById(postId)

	if err != nil {
		return
	}

	post.LikesNumber += 1

	_, err = Engine.Id(postId).Update(post)

	return
}

func (user *User) PostUnlike(postId int64) (err error) {
	if err = Engine.Ping(); err != nil {
		return
	}

	_, err = Engine.Where("post_id = ? and user_id = ?", postId, user.Id).Delete(&UserLikePost{})

	err, post := PostById(postId)

	if err != nil {
		return
	}

	post.LikesNumber -= 1

	_, err = Engine.Id(postId).Update(post)

	return
}

func PostIsLiked(postId, userId int64) (isLiked bool, err error){

	if err = Engine.Ping(); err != nil {
		return
	}

	if err = Engine.Sync2(&UserLikePost{}); err != nil {
		return
	}

	isLiked, err = Engine.Where("post_id = ? and user_id = ?", postId, userId).Get(&UserLikePost{})

	return
}

func PostById(id int64) (err error, post Post) {
	if err = Engine.Ping(); err != nil {
		return
	}

	_, err = Engine.Where("id = ?", id).Get(&post)
	return
}