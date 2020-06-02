package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type DBConnection interface {
	Open() (*gorm.DB, error)
	Close()
}

type dbConnection struct {
	conn *gorm.DB
}

func NewDBConnection() DBConnection {
	return dbConnection{}
}

func (db dbConnection) Open() (*gorm.DB, error) {
	sqlport := os.Getenv("SQLPORT")
	if sqlport == "" {
		sqlport = "5432"
	}
	sqlhost := os.Getenv("SQLHOST")
	if sqlhost == "" {
		sqlhost = "mars-colony-db.postgres.database.azure.com"
	}
	sqldb := os.Getenv("SQLDB")
	if sqldb == "" {
		sqldb = "postgres"
	}
	sqluser := os.Getenv("SQLUSER")
	if sqluser == "" {
		sqluser = ""
	}
	sqlpassword := os.Getenv("SQLPASSWORD")
	if sqlpassword == "" {
		sqlpassword = ""
	}
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", sqlhost, sqluser, sqlpassword, sqldb)
	conn, err := gorm.Open("postgres", dbURI)
	if err == nil {
		db.conn = conn
	}
	return conn, err
}

func (db dbConnection) Close() {
	db.conn.Close()
}
