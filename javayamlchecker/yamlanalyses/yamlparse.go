package yamlanalyses

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func yamlParse(yamlPath string) (flatData map[string]interface{}) {
	m := yamlFileRead(yamlPath)
	// fmt.Println(m)
	var key string
	flatData = make(map[string]interface{})
	mapFlat(m, key, &flatData)
	return
}

func mapFlat(in *map[string]interface{}, key string, flatData *map[string]interface{}) {
	for k, v := range *in {
		switch v := v.(type) {
		case string, int, bool, float32, float64, nil:
			(*flatData)[key+k] = v
			continue
		case map[string]interface{}:
			mapFlat(&v, (key + k + "."), flatData)
		default:
			fmt.Println("type error")
		}
	}
}

func yamlFileRead(yamlPath string) *map[string]interface{} {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		fmt.Println("read yaml file failed: ", err)
		os.Exit(-1)
	}
	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(-1)
	}

	return &m
}
