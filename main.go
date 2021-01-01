package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type goodHandler struct {
		msg   chan map[string]interface{}
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT が設定されていません")
	}

	good := &goodHandler{msg: make(chan map[string]interface{}, 10),}
	job	 := &jobTicker{good: good}
	go job.run()

	http.Handle("/good", good)
	if err := http.ListenAndServe(":" + port, nil); err != nil{
		log.Fatal(err)
	}
}

func (c *goodHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pass := os.Getenv("API_PASS")

		json := c.json_parse(r)
		if pass == json["Pass"] {
			c.msg <- json
		}else {
			fmt.Println("passwordが間違っています: ", json["Pass"])
		}

	}

}
