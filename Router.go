package main

import routing "github.com/qiangxue/fasthttp-routing"

func CreateRouter() *routing.Router{

	router := routing.New()

	router.Post("/<secret>/todos", Secure(CreateTodo))
	router.Get("/<secret>/todos", Secure(ReadTodos))
	router.Put("/<secret>/todos", Secure(UpdateTodo))
	router.Delete("/<secret>/todos/<id>", Secure(DeleteTodo))

	return router
}