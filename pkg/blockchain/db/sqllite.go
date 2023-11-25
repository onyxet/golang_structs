package blockchain

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"log"
)

type Database struct {
	Connection *sql.DB
}

func (d *Database) CreateDB() {
	// connect
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	d.Connection = db
}

func (d *Database) Init() {
	// create table with blockchain
	_, err := d.Connection.Exec("CREATE TABLE IF NOT EXISTS blockchain (id INTEGER PRIMARY KEY AUTOINCREMENT,timestamp DATETIME,data TEXT,prev_hash TEXT,hash TEXT);")
	if err != nil {
		panic(err)
	}

	// create table with peers
	// TBD
}
