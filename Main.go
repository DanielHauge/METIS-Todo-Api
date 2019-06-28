package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

var json = jsoniter.ConfigFastest
var env = os.Getenv("ENV")

func main() {

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", os.ModePerm)
	}

	StartProduction()

}


func StartProduction() {
	router := CreateRouter()

	srs := &fasthttp.Server{
		Handler:      router.Handler,
		ErrorHandler: HandleError,
	}
	log.Println("Server is running on localhost:443")
	log.Fatal(srs.ListenAndServe(":http"))
}
