package main

import (
	"os"
	"yamlchecker/yamlanalyses"
)

func main() {
	args := os.Args[1:3]
	yamlanalyses.YamlCheck(args[0], args[1])
}
