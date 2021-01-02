package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type goodHandler struct {
		msg   []map[string]interface{}
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT が設定されていません")
	}

	good := &goodHandler{}

	http.Handle("/good", good)

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("OK")
		if r.Method == "POST" {
			fmt.Println("ok")
			good.outputMessage()
		}
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil{
		log.Fatal(err)
	}
}

// good
func (g *goodHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pass := os.Getenv("API_PASS")

		json := g.json_parse(r)
		if pass == json["Pass"] {
			g.msg = append(g.msg, json)
			fmt.Println(g.msg)
		}else {
			fmt.Println("passwordが間違っています: ", json["Pass"])
		}

	}
}
