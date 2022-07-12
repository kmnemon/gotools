package yamlanalyses

import "testing"

func TestYamlCheck(t *testing.T) {
	yaml1 := "../testdata/1.yml"
	yaml2 := "../testdata/1.yml"

	result := YamlCheck(yaml1, yaml2)

	if result != false {
		t.Error("yaml1 and yaml2 is equal, have problems")
	}
}

func TestYamlCheckNoProblem(t *testing.T) {
	yaml1 := "../testdata/1.yml"
	yaml2 := "../testdata/2.yml"

	result := YamlCheck(yaml1, yaml2)

	if result != true {
		t.Error("yaml1 and yaml2 have no problems")
	}
}

func TestYamlCheckMissingKey(t *testing.T) {
	yaml1 := "../testdata/1.yml"
	yaml2 := "../testdata/3.yml"

	result := YamlCheck(yaml1, yaml2)

	if result != false {
		t.Error("yaml1 and yaml2 is missing key")
	}
}

func TestYamlCheckNullValue(t *testing.T) {
	//both contain null value
	yaml1 := "../testdata/4.yml"
	yaml2 := "../testdata/5.yml"

	result := YamlCheck(yaml1, yaml2)

	if result != false {
		t.Error("yaml1 and yaml2 contain null values is problems")
	}

}
