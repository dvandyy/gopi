package database

import (
	"database/sql"
	"fmt"

	"github.com/dvandyy/gopi/config"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

var db *DB

func InitDatabase() *DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Get().DB_Host,
		config.Get().DB_Port,
		config.Get().DB_User,
		config.Get().DB_Password,
		config.Get().DB_Name)

	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return &DB{Conn: conn}
}

func CreateDatabase() {
	if db == nil {
		db = InitDatabase()
		fmt.Printf("Database initialized\n")
	}
}

func GetDatabase() *DB {
	if db == nil {
		CreateDatabase()
		return db
	}
	return db
}
