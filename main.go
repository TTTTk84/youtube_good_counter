package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {
	fmt.Println("デプロイテスト")

	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil{
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>hello world</h1>"))
}
