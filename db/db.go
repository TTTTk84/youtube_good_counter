package db

import (
	"database/sql"
	"fmt"
	"os"
)


type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

func NewDB() *sql.DB {

	db := &DB{
		Host: os.Getenv("DB_HOSTNAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
	db.Username,db.Password,db.Host,db.Port,db.DBName)

	DBConn, err := sql.Open("mysql",dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DBConn.Ping(); err != nil {
		panic(err)
	}

	cmd := `CREATE TABLE IF NOT EXISTS watchtables(
						title STRING,
						url		STRING,
						likedAt STRING)`

	_, err = DBConn.Exec(cmd)
	if err != nil {
		panic(err)
	}


	return DBConn
}
