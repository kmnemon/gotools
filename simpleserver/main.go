package main

import (
	"fmt"
	"net/http"
)

func apiInfo(w http.ResponseWriter, r *http.Request) {
	// len := r.ContentLength
	// body := make([]byte, len)

	fmt.Fprintln(w, r.URL.Query())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:1234",
	}

	http.HandleFunc("/api-info", apiInfo)

	server.ListenAndServe()

}
