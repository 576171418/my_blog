package controllers

import (
	"encoding/json"
	"log"
	. "my_blog/auth"
	. "my_blog/models"
	"net/http"
	"time"
)

func CreatePost(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var response []byte

	tokenStr := request.Header.Get("token")
	user_id, succ := GetUserIdFromToken(tokenStr)

	if !succ {
		log.Fatalf("获取 UserId 失败")
	}

	err, user := UserById(user_id)

	if err != nil {
		log.Fatalf("获取用户信息失败")
	}

	title := request.Form.Get("title")
	content := request.Form.Get("content")
	var post  = Post{
		Id:             0,
		Title:          title,
		Content:        content,
		CreatedAt:      time.Now(),
		Creator:        "",
		LikesNumber:    0,
		CommentsNumber: 0,
		UserId:         0,
	}
	err = user.CreatePost(post)

	if err != nil {
		log.Fatalf("创建失败，请重试")
	}

	responseInfo := Response{
		Code:    0,
		Message: "创建成功",
		Data:    nil,
	}
	response, _ = json.Marshal(responseInfo)

	writer.Write(response)
}