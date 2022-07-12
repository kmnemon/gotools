package yamlanalyses

import (
	"testing"
)

func TestYamlWhiteListFilter(t *testing.T) {
	flatData := yamlParse("../testdata/6.yml")

	yamlFilter(&flatData)

	if len(flatData) != 2 {
		t.Error("white list filter should have 2 elements")
	}
}
