package main

import (
	"github.com/julienschmidt/httprouter"
)

// newRouter is used to set-up router based on HTTPRoutes struct
func newRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range httpRoutes {
		router.Handle(route.HTTPMethod, route.Pattern, route.HadleFunc)
	}
	return router
}
