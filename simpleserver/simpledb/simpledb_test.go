package simpledb

import (
	"bytes"
	"simpleserver/shodan"
	"testing"
)

func TestFind(t *testing.T) {
	var api shodan.APIInfo
	info, ok := Find("./db2", api, "other")
	if !ok {
		t.Error("Find failed")
	}

	apiInfo := shodan.APIInfo{QueryCredits: 33, ScanCredits: 2, Telnet: true, Plan: "org", HTTPS: false, Unlocked: true}

	if info != apiInfo {
		t.Error("api wrong")
	}

}

func TestFindByPrimaryKey(t *testing.T) {
	s, ok := FindByPrimaryKey("./db2", "other")
	if ok != true {
		t.Error("can not find  by primary key was wrong")
	}

	rs := []byte(`{"query_credits": 33, "scan_credits": 2, "telnet": true, "plan": "org", "https": false, "unlocked": true}`)
	if !bytes.Equal(s, rs) {
		t.Error("find wrong string")
	}
}

func TestFindByPrimaryKeyIfNotFound(t *testing.T) {
	s, ok := FindByPrimaryKey("./db2", "haha")
	if ok != false {
		t.Error("can not find  by primary key was wrong")
	}

	if len(s) != 0 {
		t.Error("find wrong bytes")
	}
}

func TestSplitBytesBySubbytes(t *testing.T) {
	bs := []byte(`{"keke" : {"query_credits": 56, "scan_credits": 0, "telnet": true, "plan": "edu", "https": true, "unlocked": true}}`)
	brs := []byte(`{"query_credits": 56, "scan_credits": 0, "telnet": true, "plan": "edu", "https": true, "unlocked": true}`)

	bstr, ok := splitBytesBySubbytes(bs, []byte(":"))

	if ok != true {
		t.Error("can not find substring by primary key was wrong")
	}

	if !bytes.Equal(bstr, brs) {
		t.Error("find wrong substring")
	}

}
