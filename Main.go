package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
)


var json = jsoniter.ConfigFastest
var secretGlobal = os.Getenv("secret")

func main() {

	StartLocalhost()

}


func StartLocalhost() {
	router := CreateRouter()

	srs := &fasthttp.Server{
		Handler: router.Handler,
		ErrorHandler:HandleError,
	}
	log.Println("Server is running on localhost:8080")
	log.Fatal(srs.ListenAndServeTLS(":https", "hej", "hej"))
}

func HandleError(ctx *fasthttp.RequestCtx, err error){
	ctx.Response.SetBody([]byte(err.Error()))
}

func StartProduction(handler http.Handler) {
	log.Fatal(http.ListenAndServe(":http", handler))
}
