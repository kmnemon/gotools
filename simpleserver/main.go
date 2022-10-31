package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// . "simpleserver/controller"

func foo[T any](entity T) T {
	return entity

}

type AAA struct {
	a int
	b string
}

func main() {
	FindByPrimaryKey("other", "./simpledb/db2")

	// server := http.Server{
	// 	Addr: "127.0.0.1:1234",
	// }

	// http.HandleFunc("/api-info", ApiInfo)
	// http.HandleFunc("/host/search", Search)

	// server.ListenAndServe()

}



