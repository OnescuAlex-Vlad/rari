package models

import (
	"database/sql"
	"sync"
	"fmt"

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

type DBConnection struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Mutex    sync.Mutex
	DB       *sql.DB
}

const (
	HOST   = "localhost"
	PORT   = 5432
	USER   = "postgres"
	PASS   = "root"
	DBNAME = "postgres"
)

// func NewDBConnection(username, password, host, database string, port int) *DBConnection {
// 	return &DBConnection{
// 		Username: username,
// 		Password: password,
// 		Host:     host,
// 		Port:     port,
// 		Database: database,
// 	}
// }

// func (dbConn *DBConnection) CreateConnection() (*sql.DB, error) {
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

	// dbConn.DB = db

	fmt.Println("check db connection")


	return db, nil
}


// func (dbConn *DBConnection) Close() {
// 	if dbConn.DB != nil {
// 		dbConn.DB.Close()
// 	}
// }