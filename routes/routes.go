package routes

import (
	"github.com/gorilla/mux"
	. "my_blog/controllers"
	"net/http"
)

type WebRoute struct {
	Name		string
	Method		string
	Pattern		string
	HandleFunc	http.HandlerFunc
	Middleware	mux.MiddlewareFunc
}

type WebRoutes []WebRoute

var routes = WebRoutes{
	WebRoute{
		Name:       "Login",
		Method:     "POST",
		Pattern:    "/user/login",
		HandleFunc: Login,
		Middleware: nil,
	},
	WebRoute{
		Name:       "Register",
		Method:     "POST",
		Pattern:    "/user/register",
		HandleFunc: Register,
		Middleware: nil,
	},
	WebRoute{
		Name:       "CreatePost",
		Method:     "POST",
		Pattern:    "/post/create",
		HandleFunc: CreatePost,
		Middleware: nil,
	},
	WebRoute{
		Name:		"CreateComment",
		Method:		"POST",
		Pattern:	"/post/comment",
		HandleFunc:	CreateComment,
		Middleware:	nil,
	},
	WebRoute{
		Name:       "GetComments",
		Method:     "POST",
		Pattern:    "/post/getComments",
		HandleFunc: GetComments,
		Middleware: nil,
	},
	WebRoute{
		Name:       "PostLike",
		Method:     "POST",
		Pattern:    "/post/like",
		HandleFunc: PostLike,
		Middleware: nil,
	},
	WebRoute{
		Name:       "PostUnlike",
		Method:     "POST",
		Pattern:    "/post/unlike",
		HandleFunc: PostUnlike,
		Middleware: nil,
	},
}
