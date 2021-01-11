package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)




func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT が設定されていません")
	}

	http.HandleFunc("/good", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
		pass := os.Getenv("API_PASS")

		json := json_parse(r)
		if pass == json["Pass"] {
			inputGood(json)
		}else {
			fmt.Println("passwordが間違っています: ", json["Pass"])
		}
	}
	})

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		outputMessage()
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil{
		log.Fatal(err)
	}
}
