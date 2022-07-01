package yamlparse

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

var FlatData map[string]string

func YamlParse(yamlPath string) {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		fmt.Println("read yaml failed: ", err)
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	var key string
	//init global var
	FlatData = make(map[string]string)
	mapFlat(key, &m)

}

func mapFlat(key string, in *map[string]interface{}) {
	for k, v := range *in {
		switch v := v.(type) {
		case string:
			FlatData[key+"."+k] = v
			continue
		case map[string]interface{}:
			if len(key) == 0 {
				key += k
			} else {
				key += ("." + k)
			}

			mapFlat(key, &v)
		default:
			fmt.Println("unknown type,%s: %s", reflect.TypeOf(v), v)
		}
	}
}
