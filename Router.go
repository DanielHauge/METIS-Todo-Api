package main

import (
	"github.com/valyala/fasthttprouter"
)

func CreateRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()

	router.POST("/:secret/:bucket", Secure(CreateEntry))
	router.GET("/:secret/:bucket", Secure(ReadEntries))
	router.PUT("/:secret/:bucket/:id", Secure(UpdateEntry))
	router.DELETE("/:secret/:bucket/:id", Secure(DeleteEntry))

	return router
}
