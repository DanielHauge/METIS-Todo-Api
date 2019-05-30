package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)


var json = jsoniter.ConfigFastest
var secretGlobal = os.Getenv("secret")
var env = os.Getenv("ENV")

func main() {

	if env == "DEV"{
		StartLocalhost()
	} else {
		StartProduction()
	}

}


func StartLocalhost() {
	router := CreateRouter()

	srs := &fasthttp.Server{
		Handler: router.Handler,
		ErrorHandler:HandleError,
	}
	log.Println("Server is running on localhost:443")
	log.Fatal(srs.ListenAndServeTLS(":https", "MyCertificate.crt", "MyKey.key"))

}



func StartProduction() {
	router := CreateRouter()

	srs := &fasthttp.Server{
		Handler: router.Handler,
		ErrorHandler:HandleError,
	}
	log.Println("Server is running on localhost:443")
	log.Fatal(srs.ListenAndServe(":http"))
}
