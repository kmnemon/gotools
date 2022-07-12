package yamlanalyses

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func yamlFilter(flatData *map[string]interface{}) {
	data, err := os.ReadFile("./yamlwhitelist.yml")
	// data, err := os.ReadFile("../yamlwhitelist.yml")
	if err != nil {
		fmt.Println("read yaml failed: ", err)
		return
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	// fmt.Println(m)

	for k := range *flatData {
		for k2 := range m {
			if strings.Contains(k, k2) {
				delete(*flatData, k)
			}
		}
	}

}
