package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)


var json = jsoniter.ConfigFastest
var secretGlobal = os.Getenv("secret")

func main() {
	router := CreateRouter()
	todo := Todo{2, "hej", "2019-05-21", "18:00"}

	stringm, _ := json.MarshalToString(todo)

	log.Println(stringm)
	log.Println("Server is running on localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}

/*
func StartLocalhost(handler http.Handler) {
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
	}

	srv := &http.Server{
		Addr:         ":https",
		Handler:      handler,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	//Start Server

	log.Fatal(srv.ListenAndServeTLS(os.Getenv("CERT_PUBLIC"), os.Getenv("CERT_PRIVATE")))
}

func StartProduction(handler http.Handler) {
	log.Fatal(http.ListenAndServe(":http", handler))
}
 */