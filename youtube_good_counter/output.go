package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


type discord struct {
	Content string `json:"content"`
}


type watchlist struct {
	title string
	url string
	likedAt string
}

func outputMessage(){
	// 入れる文字列を作成
	record := csv_read()
	var watchlists []watchlist
	for _, item := range record{
		add := watchlist{title: item[0],url: item[1],likedAt: item[2]}
		watchlists = append(watchlists, add)
	}


	content := fmt.Sprintf("**%d月%d日のいいね数 :** %d\n",
	time.Now().Month(), time.Now().Day(), len(watchlists))

	for _, item := range watchlists {
		content += fmt.Sprintf("いいねした動画 : **%s**   %s\n", item.title,item.url)
	}

	url := os.Getenv("DISCORD_WEBHOOK")
	discord := &discord{}
	discord.Content = content

	// discordへのリクエストを作成
	req_json, _ := json.Marshal(discord)
	resp, err := http.NewRequest("POST", url, bytes.NewReader(req_json))
	err_io(err)
	resp.Header.Set("Content-Type","application/json")

	client := new(http.Client)
	_, err = client.Do(resp) // 送信
	err_io(err)

	// csvファイル空にする
	file, err := os.OpenFile("watchlist.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0755)
	if err != nil{
		log.Fatal("file open err:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	w := [][]string {}
	writer.WriteAll(w)
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal("flush err:", err)
	}
	fmt.Println("watchlist clean")

}

func inputGood(json map[string]string) {
	file, err := os.OpenFile("watchlist.csv", os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil{
		log.Fatal("file open err:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	line := []string{json["Title"],json["Url"],json["LikedAt"]}
	writer.Write(line)
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal("flush err:", err)
	}
	fmt.Println("add record ",line)
}
