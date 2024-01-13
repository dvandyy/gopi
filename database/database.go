package database

import (
	"database/sql"
	"fmt"

	"github.com/bodatomas/gopi/config"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

var db *DB

func InitDatabase() *DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DB_Host,
		config.Cfg.DB_Port,
		config.Cfg.DB_User,
		config.Cfg.DB_Password,
		config.Cfg.DB_Name)

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
