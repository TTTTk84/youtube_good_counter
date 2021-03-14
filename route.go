package main

import (
	"fmt"
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
	fmt.Println(fmt.Sprintf("%s %s", wt.Title,wt.Url))
}


func post(w http.ResponseWriter, r *http.Request){
	var err error
	discord := data.Discord{}

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
