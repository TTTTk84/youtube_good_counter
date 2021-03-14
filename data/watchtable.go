package data

import (
	"net/http"
)

type Watchtable struct {
	Title 	string	`json:"Title"`
	Url  		string	`json:"Url"`
	Pass		string  `json:"Pass"`
}

func (wt *Watchtable) CreateWatchTable(r *http.Request) (err error){
	statement := "insert into watchtables (Title, Url) values ($1, $2) returning Title, Url"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(wt.Title,wt.Url)

	return err
}

func GetAllWatchTables() (wts []Watchtable, err error){
	statement := "select Title,Url from watchtables"
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
