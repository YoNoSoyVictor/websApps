package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/models/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

type app struct {
	db            *pgxpool.Pool
	infoLog       *log.Logger
	errorLog      *log.Logger
	users         *postgres.UserModel
	posts         *postgres.PostModel
	templateCache map[string]*template.Template
}

func main() {

	//flags
	addr := flag.String("addr", ":8080", "Network address, use as following -> '-addr ip:port', leave ip blank for localhost.")
	dsn := flag.String("dsn", "user=web password=password host=localhost port=5432 dbname=ThoughtBox", "MySQL data source name")
	flag.Parse()

	//custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//dbconn
	dbpool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	defer dbpool.Close()

	//template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	//init app
	sciver := app{
		db:            dbpool,
		infoLog:       infoLog,
		errorLog:      errorLog,
		posts:         &postgres.PostModel{DB: dbpool},
		templateCache: templateCache,
	}

	//starting server
	infoLog.Println("starting server on port", *addr)
	errorLog.Fatal(http.ListenAndServe(*addr, sciver.routes()))
}
