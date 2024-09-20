package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	models data.Models
}

func main() {

	conn := connectToDB()
	if conn == nil {
		fmt.Println("Could not connect to the postgres database")
	}

	app := Config{
		DB : conn,
		models: data.New(conn),
	}
	router := chi.NewRouter()
	router.Get("/", handleHome)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println("Error occured while opening the database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error occured while pinging the database ")
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			fmt.Println("Postgres is not ready")
			counts++
		} else {
			fmt.Println("Connected to database")
			return connection
		}

		if count > 10 {
			fmt.Println("Error occured")
			return nil
		}

		time.Sleep(2 * time.Second)
		continue
	}
}
