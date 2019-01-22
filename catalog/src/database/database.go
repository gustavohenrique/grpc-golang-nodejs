package database

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	Connect() (*sqlx.DB, error)
	GetConnection() *sqlx.DB
}

type DB struct {
	Connection *sqlx.DB
	ConnStr    string
}

func NewDB(connStr string) Database {
	db := &DB{}
	db.ConnStr = connStr
	return db
}

func (db *DB) Connect() (*sqlx.DB, error) {
	var err error
	if db.Connection == nil {
		db.Connection, err = db.getConnection()
	}
	return db.Connection, err
}

func (db *DB) GetConnection() *sqlx.DB {
	conn, _ := db.Connect()
	return conn
}

func (db *DB) getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", db.ConnStr)
	if err != nil {
		log.Printf("Failed connecting to the database. URL=%s. %s", db.ConnStr, err)
		return conn, err
	}
	conn.SetMaxOpenConns(3)
	conn.SetMaxIdleConns(3)
	conn.SetConnMaxLifetime(time.Nanosecond)
	return conn, err
}
