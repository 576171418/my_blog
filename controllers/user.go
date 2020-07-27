package controllers

import (
	"encoding/json"
	"my_blog/auth"
	. "my_blog/auth"
	. "my_blog/models"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	phone := request.Form.Get("phone")
	password := request.Form.Get("password")
	var response []byte
	has, err := AuthUserPassword(phone, password)
	if err != nil {
		responseInfo := Response{
			Code:    -1,
			Message: "数据库查询失败，请重试",
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	} else {
		if has {
			token, _ := auth.GenerateToken(1)
			responseInfo := Response{
				Code:    1,
				Message: "登录成功",
				Data: JwtToken{Token:token},
			}
			response, _ = json.Marshal(responseInfo)
		} else {
			responseInfo := Response{
				Code:    -1,
				Message: "手机号或密码错误",
				Data:    nil,
			}
			response, _ = json.Marshal(responseInfo)
		}
	}
	writer.Write(response)
}

func Register(writer http.ResponseWriter, request *http.Request) {
	var response []byte
	request.ParseForm()
	errInfo := ""
	phone := request.Form.Get("phone")
	password := request.Form.Get("password")
	name := request.Form.Get("name")
	if phone == "" {
		errInfo = "手机号为空"
	}
	if CheckPhoneExist(phone) {
		errInfo = "此手机号已存在"
	}
	if password == "" {
		errInfo = "密码不能为空"
	}
	if name == "" {
		name = phone
	}
	user := User{
		Name:      name,
		Phone:     phone,
		CreatedAt: time.Time{},
		Password:  password,
	}
	if errInfo != "" {
		responseInfo := Response{
			Code:    -1,
			Message: errInfo,
			Data:    nil,
		}
		response, _ = json.Marshal(responseInfo)
	} else {
		err := user.Create()
		if err != nil {
			errInfo = "连接数据库失败，请重试"
			responseInfo := Response{
				Code:    -1,
				Message: errInfo,
				Data:    nil,
			}
			response, _ = json.Marshal(responseInfo)
		} else {
			responseInfo := Response{
				Code:    0,
				Message: "注册成功",
				Data:    nil,
			}
			response, _ = json.Marshal(responseInfo)
		}
	}
	writer.Write(response)
}