package controller

import (
	"fmt"
	"net/http"
	"simpleserver/service"
)

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(403)
		fmt.Fprintf(w, "invalid key: %q\n", key)
		return
	}

	data, ok := service.FindByPrimaryKeyService(key)
	if !ok {
		w.WriteHeader(403)
		fmt.Fprintf(w, "invalid key: %q\n", key)
		return
	}

	w.Write(data)
}

func Search(w http.ResponseWriter, r *http.Request) {

}
