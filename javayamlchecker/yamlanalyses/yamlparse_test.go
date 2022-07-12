package yamlanalyses

import (
	"reflect"
	"testing"
)

func TestYamlParse(t *testing.T) {
	flatData := yamlParse("../testdata/1.yml")

	flatDataExpect := map[string]interface{}{
		"schedule.admin.address":     "http://123456",
		"schedule.executor.appname":  "epms-fss-test",
		"schedule.executor.appname2": "epms-fss-test2",
	}

	if !reflect.DeepEqual(flatDataExpect, flatData) {
		t.Error("flatData and flatDataExpect should be identical")
	}

}

func TestYamlFileRead(t *testing.T) {
	data := yamlFileRead("../testdata/1.yml")
	if data == nil {
		t.Error("yamlfile read have something wrong")
	}

}
