package main

import (
	"log"
	"net/http"
)

const (
	host = "127.0.0.1"
	port = "4000"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/table", showTable)
	mux.HandleFunc("/table/insert", insertTable)

	log.Println("Server address: http://" + host + ":" + port)
	err := http.ListenAndServe(host+":"+port, mux)
	log.Fatal(err)
}
