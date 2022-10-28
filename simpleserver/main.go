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

func FindByPrimaryKey(key string, db string) (string, bool) {
	f, err := os.Open(db)
	if err != nil {
		fmt.Println("err to open file")
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var line []byte
	var str string
	var ok bool

	for err == nil && !ok {
		line, _, err = r.ReadLine()
		str, ok = findStringByPrimaryKey(key, string(line))
	}

	return str, ok
}

func findStringByPrimaryKey(key string, line string) (string, bool) {
	i := strings.Index(":", line)
	fmt.Println(i)
	fmt.Println(line)

	return line, false
}
