package main

import (
	"errors"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)


func Secure(handlefunction func (context *routing.Context) error ) func(context *routing.Context) error{
	return func(context *routing.Context) error{
		secret := context.Param("secret")
		context.SetContentType("application/json")
		if secret != secretGlobal{
			context.SetStatusCode(fasthttp.StatusUnauthorized)
			return errors.New("Unauthorized access")
		} else {
			err := handlefunction(context)
			if err != nil{ log.Println(err) }
			return err
		}
	}
}

func CreateTodo(context *routing.Context) error{

	var todo Todo
	body := context.PostBody()

	err := json.Unmarshal(body, &todo)
	if err != nil { return err }

	err = Create(todo)
	if err != nil { return err }

	context.SetStatusCode(fasthttp.StatusOK)
	return err // will be nil
}

func ReadTodos(context *routing.Context) error{
	todos, err := Read()
	if err != nil { return err }
	bytes, err := json.Marshal(todos)
	if err != nil { return err }
	context.SetBody(bytes)
	context.SetStatusCode(fasthttp.StatusOK)
	return err // will be nil
}

func UpdateTodo(context *routing.Context) error{

	var todo Todo
	err := json.Unmarshal(context.PostBody(), &todo)
	if err != nil { return err }

	err = Update(todo)
	if err != nil { return err }

	context.SetStatusCode(fasthttp.StatusOK)

	return err // will be nil
}

func DeleteTodo(context *routing.Context) error{
	stringId := context.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil { return err }

	err = Delete(id)
	if err != nil { return err }

	context.SetStatusCode(fasthttp.StatusOK)

	return err // will be nil
}
