package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func err_io(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}


func (c *goodHandler) json_parse(r *http.Request) map[string]interface{}{
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	err_io(err)

	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	err_io(err)

	return jsonBody
}


type discord struct {
	Content string `json:"content"`
}

func (g *goodHandler) outputMessage() {
	// 入れる文字列を作成
	content := fmt.Sprintf("**%d月%d日のいいね数 :** %d\n",
	time.Now().Month(), time.Now().Day(), len(g.msg))
	for _, mes := range g.msg {
		content += fmt.Sprintf("いいねした動画 : **%s**   %s\n", mes["Title"], mes["Url"])
	}
	fmt.Println(content)


	url := os.Getenv("DISCORD_WEBHOOK")
	discord := &discord{}
	discord.Content = content

	// リクエストを作成
	req_json, _ := json.Marshal(discord)
	resp, err := http.NewRequest("POST", url, bytes.NewReader(req_json))
	err_io(err)
	resp.Header.Set("Content-Type","application/json")

	client := new(http.Client)
	_, err = client.Do(resp) // 送信
	err_io(err)

	g.msg = nil
}
