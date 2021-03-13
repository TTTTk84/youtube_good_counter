package main

import (
	"log"
	"net/http"

	"github.com/TTTTk84/youtube_good_counter/data"
)

func good(w http.ResponseWriter, r *http.Request){
	wt := data.Watchtable{}
	err := wt.CreateWatchTable(r)
	if err != nil{
		log.Fatal(err)
	}
}


func post(w http.ResponseWriter, r *http.Request){
	discord := data.Discord{}
	err := discord.PostWebhook()
	if err != nil {
		log.Fatal(err)
	}

	err = data.DeleteAllWatchTables()
	if err != nil {
		log.Fatal(err)
	}

}
