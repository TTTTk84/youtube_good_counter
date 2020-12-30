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
		log.Fatal("$PORT must be set")
	}
	fmt.Println("デプロイテスト")

	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":" + port, nil); err != nil{
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>hello world</h1>"))
}
