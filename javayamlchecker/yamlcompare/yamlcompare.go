package yamlcompare

import "fmt"

func YamlCompare(yamlPath1 string, yamlPath2 string, a, b map[string]string) {
	if len(a) != len(b) {
		fmt.Println("two compared yaml file have different nummber of lines")
	}
	for k, v := range a {
		if b[k] == "" {
			fmt.Println(yamlPath2 + "don't have the para: " + k)
		} else if v == b[k] {
			fmt.Println()
		}
	}

}
