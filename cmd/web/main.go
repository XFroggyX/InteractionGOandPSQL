package main

import (
	"log"
	"net/http"
	"path/filepath"
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

	fileServer := http.FileServer(customizableFileSystem{http.Dir("./static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server address: http://" + host + ":" + port)
	err := http.ListenAndServe(host+":"+port, mux)
	log.Fatal(err)
}

type customizableFileSystem struct {
	fs http.FileSystem
}

func (cfs customizableFileSystem) Open(path string) (http.File, error) {
	f, err := cfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "page.html")
		if _, err := cfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
