package controllers

import (
	"encoding/json"
	"fmt"
	. "my_blog/auth"
	. "my_blog/models"
	"net/http"
	"strconv"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var response []byte

	tokenStr := r.Header.Get("token")
	user_id, succ := GetUserIdFromToken(tokenStr)

	if !succ {
		panic(fmt.Sprint("无效的 Token"))
	}

	err, user := UserById(user_id)

	if err != nil {
		panic(fmt.Sprint("获取用户信息失败"))
	}

	post_id := r.Form.Get("post_id")
	postId, err := strconv.ParseInt(post_id, 10, 64)
	content := r.Form.Get("content")

	err = user.CreateComment(postId, content)

	if err != nil {
		panic(fmt.Sprint("创建评论失败，请重试"))
	}

	responseInfo := Response{
		Code:    0,
		Message: "评论成功",
		Data:    nil,
	}

	response, _ = json.Marshal(responseInfo)

	w.Write(response)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var response []byte

	post_id := r.Form.Get("post_id")
	postId, err := strconv.ParseInt(post_id, 10, 64)

	if err != nil {
		panic(fmt.Sprint("post_id 不存在"))
	}

	comments, err := GetCommentsFromPostId(postId)

	if err != nil {
		panic(fmt.Sprint("获取评论列表失败"))
	}

	responseInfo := Response{
		Code:    0,
		Message: "操作成功",
		Data:    comments,
	}

	response, _ = json.Marshal(responseInfo)

	w.Write(response)
}