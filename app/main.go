package main

import (
	"log"

	"github.com/TTTTk84/youtube_good_counter/db"
	"github.com/TTTTk84/youtube_good_counter/youtube_GC/http"
	"github.com/TTTTk84/youtube_good_counter/youtube_GC/mysql"
	"github.com/TTTTk84/youtube_good_counter/youtube_GC/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err: .env loading ", err)
	}
}


func main() {
	env_load()
	router := gin.Default()

	GC_sqlRepo := mysql.NewGC_mysql(db.NewDB())
	GC_usecase := usecase.NewGCUsecase(GC_sqlRepo)
	http.NewGC_handler(router, GC_usecase)

	router.Run(":8080")
}
