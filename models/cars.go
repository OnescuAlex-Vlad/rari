package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Car struct {
	Id        int    `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Engine    string `json:"engine"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

const (
	HOST   = "localhost"
	PORT   = 5432
	USER   = "postgres"
	PASS   = "root"
	DBNAME = "postgres"
)

func CreateConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASS, DBNAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	return db, nil
}

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Your code using the database connection goes here.
}
