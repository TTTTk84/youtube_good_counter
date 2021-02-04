package main

import (
	"log"

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

}
