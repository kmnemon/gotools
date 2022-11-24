package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get("http://127.0.0.1:1234/api")
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)

	}

}
