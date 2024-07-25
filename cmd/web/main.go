package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.elmm.net/internal/models"
)

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()
	db, err := openDB(*dsn)
	app := &Application{
		Logger:   slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		snippets: &models.SnippetModel{DB: db},
	}

	app.Logger.Info("Starting server", slog.Any("addr", *addr))
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
	app.Logger.Info(fmt.Sprintf("%v", db.Ping()))
	defer db.Close()

	error := http.ListenAndServe(fmt.Sprintf(":%v", *addr), app.router())
	app.Logger.Error(error.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
