package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TTTTk84/youtube_good_counter/data"
)

func good(w http.ResponseWriter, r *http.Request){
	wt := data.Watchtable{}

	err := data.JsonParse(r, &wt)
	if err != nil {
		log.Fatal(err)
	}

	if !data.PassCheck(wt.Pass) {
		log.Fatal(err)
	}

	err = wt.CreateWatchTable(r)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%s %s", wt.Title,wt.Url))
}


func post(w http.ResponseWriter, r *http.Request){
	discord := data.Discord{}
	var pass map[string]string

	err := data.JsonParse(r, &pass)
	if err != nil {
		log.Fatal(err)
	}

	if !data.PassCheck(pass["Pass"]) {
		log.Fatal(err)
	}

	err = discord.PostWebhook()
	if err != nil {
		log.Fatal(err)
	}

	err = data.DeleteAllWatchTables()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE ALL watchtables")

}
