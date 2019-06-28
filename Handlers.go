package main

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"strconv"
)

func Filter(handlefunction fasthttprouter.Handle) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		ctx.SetContentType("application/json")
		handlefunction(ctx, ps)
	})
}

func CreateEntry(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {

	var data interface{}
	body := ctx.PostBody()

	err := json.Unmarshal(body, &data)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	bucket := ps.ByName("bucket")
	id, err := Create(data, bucket)

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody([]byte(strconv.Itoa(id)))
	ctx.SetContentType("text/plain")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func ReadEntries(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	bucket := ps.ByName("bucket")
	entries, err := Read(bucket)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(entries)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(bytes)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func UpdateEntry(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	var data interface{}
	err = json.Unmarshal(ctx.PostBody(), &data)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	entry := Entry{Data: data, Id: id}
	bucket := ps.ByName("bucket")
	err = Update(entry, bucket)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func DeleteEntry(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	bucket := ps.ByName("bucket")
	err = Delete(id, bucket)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func HandleError(ctx *fasthttp.RequestCtx, err error) {
	ctx.Response.SetBody([]byte(err.Error()))
}
