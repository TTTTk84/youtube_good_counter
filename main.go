package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TTTTk84/youtube_good_counter/youtube_good_counter"
	"github.com/joho/godotenv"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err: .env loading ", err)
	}
}


func main() {
	env_load()
	youtube_good_counter.Db_init()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT が設定されていません")
	}

	http.HandleFunc("/good", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			pass := os.Getenv("API_PASS")

			json := youtube_good_counter.Json_parse(r)
			if pass == json["Pass"] {
				youtube_good_counter.CreateGoodColumn(json)
			}else {
				fmt.Println("passwordが間違っています: ", json["Pass"])
			}
	  }
	})

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		//youtube_good_counter.OutputMessage()
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil{
		log.Fatal(err)
	}
}
