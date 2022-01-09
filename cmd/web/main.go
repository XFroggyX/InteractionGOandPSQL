package main

import (
	"context"
	"flag"
	postgresql "github.com/XFroggyX/InteractionGOandPSQL/pkg/models/postgre"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	ctx        context.Context
	listTables map[string]interface{}
}

func main() {
	port := flag.String("port", "4000", "Сетевой порт")
	host := flag.String("addr", "127.0.0.1", "Сетевой адрес")
	dbURL := flag.String("db", "postgres://admin:admin@localhost:5432/Countries",
		"Название Postgresql базы данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.Background()
	db, err := openDB(ctx, *dbURL)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	listTables := make(map[string]interface{})

	listTables["Countries"] = postgresql.CountriesModel{DB: db}
	listTables["Languages"] = postgresql.LanguagesModel{DB: db}
	listTables["GovernmentForms"] = postgresql.GovernmentFormsModel{DB: db}
	listTables["TerritorySizes"] = postgresql.TerritorySizesModel{DB: db}
	listTables["Religions"] = postgresql.ReligionsModel{DB: db}
	listTables["Associations"] = postgresql.AssociationsModel{DB: db}
	listTables["AssociationsOfCountries"] = postgresql.AssociationsOfCountriesModel{DB: db}
	listTables["Сontinents"] = postgresql.СontinentsModel{DB: db}
	listTables["СontinentsOfCountries"] = postgresql.СontinentsOfCountriesModel{DB: db}

	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		ctx:        ctx,
		listTables: listTables,
	}

	srv := &http.Server{
		Addr:     *host + ":" + *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Server address: http://" + *host + ":" + *port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
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
