package testserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = 41414
)

func RunTestServer() {
	http.HandleFunc("/", handlerRoot)
	Serve()
}

func Serve() {
	addr := fmt.Sprintf("%s:%d", HOST, PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Could not listen to %s: %+v", addr, err)
	}

	err = http.Serve(listener, nil)
	log.Printf("Started server: %+v", err)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	SendJson(w, http.StatusOK, `{"hello": "world"}`)
}

func SendJson(w http.ResponseWriter, statusCode int, content string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(content))
}
