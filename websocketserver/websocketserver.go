package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{ //2
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	wsAddr     string  
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "address for websocket connection")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles("logger.js") //3
	if err != nil {
		panic(err)
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) //4
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage() //5
		if err != nil {
			return
		}
		fmt.Printf("From %s:%s\n", conn.RemoteAddr().String(), string(msg)) //6
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript") //7
	jsTemplate.Execute(w, wsAddr)                            //8
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)     //9
	r.HandleFunc("/k.js", serveFile) //10
	log.Fatal(http.ListenAndServe(":8080", r))
}
