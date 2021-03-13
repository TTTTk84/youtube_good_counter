package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Watchtable struct {
	Title 	string
	User  	string
	LikedAt	string
}

func (wt *Watchtable) CreateWatchTable(r *http.Request) (err error){
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))

	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	err = json.Unmarshal(body[:length], &wt)

	return err
}
