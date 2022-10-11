package main

import (
	"net/http"
	. "simpleserver/controller"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:1234",
	}

	http.HandleFunc("/api-info", ApiInfo)

	server.ListenAndServe()

}
