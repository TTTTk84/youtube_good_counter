package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Db *sql.DB

func SqlConnect(){
	var err error
	sourceName := fmt.Sprintf("dbname=%s sslmode=disable user=%s password=%s",
										os.Getenv("DB_NAME"),os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"))
	Db, err = sql.Open("postgres", sourceName)
	if err != nil{
		log.Fatal(err)
	}
	return
}
