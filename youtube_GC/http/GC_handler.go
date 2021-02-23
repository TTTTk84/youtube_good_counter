package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/TTTTk84/youtube_good_counter/domain"
	"github.com/gin-gonic/gin"
)


type GCHandler struct {
	gcusecase domain.GCUsecase
}

type discord struct {
	Content string `json:"content"`
}

func NewGC_handler(c *gin.Engine, gu domain.GCUsecase) {
	handler := &GCHandler{
		gcusecase: gu,
	}

	c.POST("/GoodCounte", handler.AddGC)
	c.POST("/GCOnceday", handler.GCOnceday)
}

func (h *GCHandler) AddGC(c *gin.Context) {
	wt, err := h.gcusecase.JsonParse(c.Request)
	if err != nil {
		log.Fatal("err: jsonparse ", err)
	}

	err = h.gcusecase.AddGC(wt)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *GCHandler) GCOnceday(c *gin.Context) {
	// wt, err := h.gcusecase.GetAll()
	// if err != nil {
	// 	log.Fatal("err: ", err)
	// }

	content := "test"

	discord := &discord{}
	url := os.Getenv("DISCORD_WEBHOOK")
	discord.Content = content

	json_byte, _ := json.Marshal(discord)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(json_byte))
	if err != nil {
		log.Fatal("err:", err)
	}

	defer res.Body.Close()



}
