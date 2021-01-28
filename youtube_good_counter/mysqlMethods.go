package youtube_good_counter

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(DBMS, DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	return db
}
