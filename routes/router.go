package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.HandleFunc))
		} else {
			r.Handler(route.HandleFunc)
		}
	}
	return router
}