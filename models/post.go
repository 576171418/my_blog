package models

import "time"

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