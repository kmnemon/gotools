package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleserver/service"
	. "simpleserver/shodan"
)

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("key")
	if value != "keke" {
		w.WriteHeader(403)
		fmt.Fprintf(w, "invalid key: %q\n", value)
		return
	}

	var apiInfo APIInfo = *service.ApiInfoService()

	apiInfoJson, err := json.MarshalIndent(&apiInfo, "", "\t")
	if err != nil {
		fmt.Println("Failed to marshal APIInfo: ", err)
		return
	}

	w.Write(apiInfoJson)
}

func Search(w http.ResponseWriter, r *http.Request) {

}
