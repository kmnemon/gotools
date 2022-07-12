package yamlanalyses

import (
	"strings"
)

func yamlFilter(flatData *map[string]interface{}) {
	yamlPath := "./yamlwhitelist.yml"
	// yamlPath := "../yamlwhitelist.yml"

	m := yamlFileRead(yamlPath)

	for k := range *flatData {
		for k2 := range *m {
			if strings.Contains(k, k2) {
				delete(*flatData, k)
			}
		}
	}

}
