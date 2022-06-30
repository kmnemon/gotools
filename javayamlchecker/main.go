package main

import (
	"fmt"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

var data = `
schedule:
  admin:
    address: http://123456
  executor:
    appname: epms-fss-test
`

func yamlParse(data string) {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	mapFlat(m)

}

func mapFlat(in map[interface{}]interface{}) map[string]string {
	for k, v := range in {

	}
}

var flatData map[string]string

func mapFlatttt(in map[interface{}]interface{}) interface{} {
	for _, v := range in {
		switch v := v.(type) {
		case string:
			return v
		case map[interface{}]interface{}:
			mapFlatttt(v)
		default:
			fmt.Println("unknown type,%s: %s", reflect.TypeOf(v), v)
		}
	}
}

func main() {
	yamlparse(data)

}
