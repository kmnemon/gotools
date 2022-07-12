package main

import (
	"fmt"
	"os"
	"yamlchecker/yamlanalyses"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "help" {
		fmt.Println(
			`yamlChecker is a tool for check yaml configuration files.

Usage:
	./yamlchecker <yaml1 file path> <yaml2 file path>

Check Scope:
	1. A == B
	2. A have some configure item that B don't have and vice visa
	3. Both A and B have some null values

	to report bugs, please contact the author: 
	liuke978

		`)
	} else if len(args) == 2 {
		yamlanalyses.YamlCheck(args[0], args[1])
	} else {
		fmt.Println(
			`unknown command
Run './yamlchecker help' for usage.`)
	}

}
