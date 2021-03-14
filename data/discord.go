package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Discord struct {
	Content string `json:"content"`
}

func (dis *Discord) PostWebhook() (err error){
	wts, err := GetAllWatchTables()
	if err != nil {
		return err
	}

	content := fmt.Sprintf("**%d月%d日のいいね数 :** %d\n",
	time.Now().Month(), time.Now().Day(), len(wts))

	for _, item := range wts {
		content += fmt.Sprintf("いいねした動画 : **%s**   %s\n", item.Title,item.Url)
	}

	url := os.Getenv("DISCORD_WEBHOOK")
	dis.Content = content

	req_json, err := json.Marshal(dis)
	if err != nil{
		return err
	}

	resp, err := http.NewRequest("POST", url, bytes.NewReader(req_json))
	if err != nil{
		return err
	}

	resp.Header.Set("Content-Type","application/json")

	client := new(http.Client)
	_, err = client.Do(resp)

	return err
}
