package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
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


func PassCheck(pass string) (bool){
	if pass != os.Getenv("API_PASS"){
		return false
	}
	return true
}


func JsonParse(r *http.Request, v interface{}) (error){
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	err := json.Unmarshal(body[:length], &v)

	return err
}
