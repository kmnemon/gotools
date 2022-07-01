package main

import (
	"fmt"
	"yamler/yamlparse"
)

func main() {
	var yamlPath1 string = "/Users/keliu/tmp/1.yaml"
	yamlparse.YamlParse(yamlPath1)
	var flatData1 map[string]string = yamlparse.FlatData

	var yamlPath2 string = "/Users/keliu/tmp/1.yaml"
	yamlparse.YamlParse(yamlPath2)
	var flatData2 map[string]string = yamlparse.FlatData

	fmt.Println(flatData1)
	fmt.Println(flatData2)
}
