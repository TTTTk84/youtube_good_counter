package main

import (
	"net/http"
	"os"

	"github.com/TTTTk84/youtube_good_counter/data"
)

func main() {
	//env_load()
	port := os.Getenv("PORT")
	data.SqlConnect()


	http.HandleFunc("/good", good)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":" + port, nil)
}
