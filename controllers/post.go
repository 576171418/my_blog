package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	. "my_blog/auth"
	. "my_blog/models"
	"net/http"
	"strconv"
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

func PostLike(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var response []byte

	tokenStr := r.Header.Get("token")
	userId, succ := GetUserIdFromToken(tokenStr)

	if !succ {
		panic(fmt.Sprint("无效的 token"))
	}

	err, user := UserById(userId)

	if err != nil {
		panic(fmt.Sprint("获取用户信息失败"))
	}

	post_id := r.Form.Get("post_id")
	postId, _ := strconv.ParseInt(post_id, 10, 64)

	isLiked, err := PostIsLiked(postId, user.Id)

	if err != nil {
		panic(fmt.Sprint("数据库查询失败"))
	}

	if !isLiked {
		_ = user.PostLike(postId)
		responseInfo := Response{
			Code:    0,
			Message: "点赞成功",
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	} else {
		responseInfo := Response{
			Code:    -1,
			Message: "已点赞过此文章",
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	}
	w.Write(response)

}

func PostUnlike(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var response []byte

	tokenStr := r.Header.Get("token")

	userId, succ := GetUserIdFromToken(tokenStr)

	if !succ {
		panic(fmt.Sprint("无效的 token"))
	}

	err, user := UserById(userId)

	if err != nil {
		panic(fmt.Sprint("获取用户信息失败"))
	}

	post_id := r.Form.Get("post_id")
	postId, _ := strconv.ParseInt(post_id, 10, 64)

	isLiked, err := PostIsLiked(postId, user.Id)

	if err != nil {
		panic(fmt.Sprint("数据库查询失败"))
	}

	if !isLiked {
		responseInfo := Response{
			Code:    -1,
			Message: "此文章没有点赞记录",
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	} else {
		_ = user.PostUnlike(postId)
		responseInfo := Response{
			Code:    0,
			Message: "取消点赞成功",
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	}
	w.Write(response)

}