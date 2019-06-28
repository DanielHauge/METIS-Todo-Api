package main

import (
	"github.com/valyala/fasthttprouter"
)

func CreateRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()

	router.POST("/:bucket", Filter(CreateEntry))
	router.GET("/:bucket", Filter(ReadEntries))
	router.PUT("/:bucket/:id", Filter(UpdateEntry))
	router.DELETE("/:bucket/:id", Filter(DeleteEntry))

	return router
}
