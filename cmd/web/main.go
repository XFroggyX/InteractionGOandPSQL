package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	port := flag.String("port", "4000", "Сетевой порт")
	host := flag.String("addr", "127.0.0.1", "Сетевой адрес")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *host + ":" + *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Server address: http://" + *host + ":" + *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
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
