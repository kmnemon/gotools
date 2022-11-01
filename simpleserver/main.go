package main

import (
	"net/http"
	"simpleserver/controller"
)

func main() {

	server := http.Server{
		Addr: "127.0.0.1:1234",
	}

	http.HandleFunc("/api-info", controller.ApiInfo)
	http.HandleFunc("/host/search", controller.Search)

	server.ListenAndServe()

}
