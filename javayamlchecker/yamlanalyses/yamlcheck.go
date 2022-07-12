package yamlanalyses

import "fmt"

func YamlCheck(yamlPath1 string, yamlPath2 string) (r bool) {
	r = true
	a := yamlParse(yamlPath1)
	b := yamlParse(yamlPath2)
	yamlFilter(&a)
	yamlFilter(&b)

	if len(a) != len(b) {
		fmt.Println("... two checked yaml files have different number of keys ...")
		r = false
	}

	for k, v := range a {
		if v == nil {
			fmt.Println("... " + yamlPath1 + ", have null value: " + k + " ...")
			r = false
		} else if _, ok := b[k]; !ok {
			fmt.Println("... " + yamlPath2 + ", don't have the key: " + k + " ...")
			r = false
		} else if a[k] == b[k] {
			fmt.Println("... key have equal value: " + k + " ...")
			r = false
		}
	}

	for k, v := range b {
		if v == nil {
			fmt.Println("... " + yamlPath2 + ", have null value: " + k + " ...")
			r = false
		} else if _, ok := a[k]; !ok {
			fmt.Println("... " + yamlPath1 + ", don't have the key: " + k + " ...")
			r = false
		}
	}

	return

}
