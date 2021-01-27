package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)



func err_io(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}


func json_parse(r *http.Request) map[string]string{
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	err_io(err)

	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	err_io(err)

	return jsonBody
}


// csvを読み込む
func csv_read() [][]string {
	file, err := os.Open("watchlist.csv")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// レコード数のチェックを行わない
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return record
}
