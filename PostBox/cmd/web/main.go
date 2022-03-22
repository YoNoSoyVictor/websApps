package main

import (
	"context"
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/models/postgres"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
)

type app struct {
	db            *pgxpool.Pool
	infoLog       *log.Logger
	errorLog      *log.Logger
	session       *sessions.Session
	users         *postgres.UserModel
	posts         *postgres.PostModel
	templateCache map[string]*template.Template
}

func main() {

	//flags
	addr := flag.String("addr", ":8080", "Network address, use as following -> '-addr ip:port', leave ip blank for localhost.")
	dsn := flag.String("dsn", "user=web password=password host=localhost port=5432 dbname=PostBox", "Postgres data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
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

	//caching templates
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	//session
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	//initializing app
	PostBox := app{
		db:            dbpool,
		infoLog:       infoLog,
		errorLog:      errorLog,
		session:       session,
		users:         &postgres.UserModel{DB: dbpool},
		posts:         &postgres.PostModel{DB: dbpool},
		templateCache: templateCache,
	}

	//initalize custom tls Config
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      PostBox.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//starting server
	infoLog.Println("starting server on port", *addr)
	errorLog.Fatal(srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"))
}
