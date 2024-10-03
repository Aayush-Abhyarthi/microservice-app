package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func validateCreds(sampledata RequestData) bool {

	fmt.Println("In the validate function")
	return true

	// This function will interact with the master database to verify the credentials. First setup the database and then implement this function to interact with it.

}

func connectDB() (*sql.DB, error) {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database!")
	return db, nil
}

func queryData(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, "John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted record with ID: %d\n", lastInsertID)
}
