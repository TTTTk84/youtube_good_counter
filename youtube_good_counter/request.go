package youtube_good_counter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)


type discord struct {
	Content string `json:"content"`
}


type Watchtables struct {
	gorm.Model
	Title string
	Url string
	LikedAt string
}

type Watchtable struct {
	Title string
	Url string
	LikedAt string
}

func OutputMessage(){
	db := sqlConnect()
	defer db.Close()

	watchtables := []Watchtable{}
	db.Find(&watchtables)

	content := fmt.Sprintf("**%d月%d日のいいね数 :** %d\n",
	time.Now().Month(), time.Now().Day(), len(watchtables))

	for _, item := range watchtables {
		content += fmt.Sprintf("いいねした動画 : **%s**   %s\n", item.Title,item.Url)
	}
	fmt.Println(content)

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

	db.Exec("DELETE FROM watchtables")

}

func CreateGoodColumn(json map[string]string) {
	db := sqlConnect()
	defer db.Close()

	watchtable := &Watchtable{Title: json["Title"], Url: json["Url"], LikedAt: json["LikedAt"]}
	err := db.Create(watchtable).Error
	if err != nil {
		fmt.Println("insert column err:", err)
		return
	}
	fmt.Println("add column:",watchtable.Title, watchtable.Url, watchtable.LikedAt)
}
