package main

import (
	"fmt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello client")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:1234",
	}

	http.HandleFunc("/api", api)

	server.ListenAndServe()

}
