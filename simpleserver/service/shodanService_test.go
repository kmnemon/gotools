package service

import (
	"bytes"
	"testing"
)

func TestFindByPrimaryKeyService(t *testing.T) {
	key := "keke"
	data, b := FindByPrimaryKeyService(key)

	if b != true {
		t.Error("find by primary key service wrong")
	}

	var targetData []byte = []byte(`{"query_credits": 56, "scan_credits": 0, "telnet": true, "plan": "edu", "https": true, "unlocked": true}`)
	if !bytes.Equal(data, targetData) {
		t.Error("find by primary key data wrong")
	}
}
