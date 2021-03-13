package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Watchtable struct {
	Title 	string
	Url  	string
	LikedAt	string
}

func (wt *Watchtable) CreateWatchTable(r *http.Request) (err error){
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))

	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	err = json.Unmarshal(body[:length], &wt)
	if err != nil {
		return err
	}

	statement := "insert into watchtables (Title, User, LikedAt) values ($1, $2, $3) Title, Url, LikedAt"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(wt.Title,wt.Url,wt.LikedAt)

	return err
}

func GetAllWatchTables() (wts []Watchtable, err error){
	statement := "select * from watchtables"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	for rows.Next() {
		var wt Watchtable
		err = rows.Scan(
			&wt.Title,
			&wt.Url,
			&wt.LikedAt,
		)
		if err != nil {
			return nil, err
		}
		wts = append(wts, wt)
	}

	return wts, err
}

func DeleteAllWatchTables() (err error){
	statement := "delete from watchtables"
	_, err = Db.Exec(statement)
	return err
}
