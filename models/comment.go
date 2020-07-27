package models

import "time"

type Comment struct {
	Id			int64		`json:"id"`
	PostId		int64		`json:"post_id" xorm:"post_id"`
	Content		string		`json:"content" xorm:"content"`
	UserId		int64		`json:"user_id" xorm:"user_id"`
	UserName	string		`json:"user_name" xorm:"user_name"`
	CreatedAt	time.Time	`json:"created_at" xorm:"created_at"`
}

func (user *User) CreateComment(postId int64, content string) (err error) {
	if err = Engine.Ping(); err != nil {
		return
	}

	if err = Engine.Sync2(&Comment{}); err != nil {
		return
	}

	_, err = Engine.Insert(&Comment{
		Id:       	0,
		PostId:   	postId,
		Content:  	content,
		UserId:   	user.Id,
		UserName: 	user.Name,
		CreatedAt:	time.Now(),
	})

	return err
}


func GetCommentsFromPostId(post_id int64) (comments []Comment, err error) {
	if err = Engine.Ping(); err != nil {
		return
	}

	if err = Engine.Sync2(&Comment{}); err != nil {
		return
	}

	err = Engine.Where("post_id = ?", post_id).Find(&comments)

	return

}