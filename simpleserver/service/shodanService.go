package service

import (
	. "simpleserver/shodan"
)
/*
#cgo CFLAGS: -fno-debug	
*/
func ApiInfoService() *APIInfo {
	return &APIInfo{
		QueryCredits: 3,
		ScanCredits:  1,
		Telnet:       true,
		Plan:         "abc",
		HTTPS:        false,
		Unlocked:     true,
	}
}
