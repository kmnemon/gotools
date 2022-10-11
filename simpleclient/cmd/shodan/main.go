package main

import (
	"fmt"
	"simpleclient/shodan"
)

func main() {
	const API_KEY = "keke"

	client := shodan.New(API_KEY)
	apiInfo, err := client.APIInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", apiInfo)

}
