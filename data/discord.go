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
	content string
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
	dis.content = content
	req_json, _ := json.Marshal(dis)
	resp, err := http.NewRequest("POST", url, bytes.NewReader(req_json))
	if err != nil{
		return err
	}

	resp.Header.Set("Content-Type","application/json")

	client := new(http.Client)
	_, err = client.Do(resp)

	return
}
