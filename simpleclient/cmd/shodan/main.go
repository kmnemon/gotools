package main

import (
	"fmt"
	"log"
	"simpleclient/shodan"
)

func main() {

	const API_KEY = "keke"

	s := shodan.New(API_KEY)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
		return
	}

	fmt.Printf("%#v", info)

	hostSearch, err := s.HostSearch("sdfsdf")
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%18d\n", host.HostNames, host.Location.PostalCode)
	}

}
