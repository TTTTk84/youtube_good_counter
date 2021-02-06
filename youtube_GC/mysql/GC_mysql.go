package mysql

import (
	"database/sql"
	"errors"

	"github.com/TTTTk84/youtube_good_counter/domain"
)
type gcsqlRepository struct {
	Conn *sql.DB
}

func NewGC_mysql(Conn *sql.DB) domain.GCsqlRepository {
	return &gcsqlRepository{Conn}
}

func (m *gcsqlRepository) Store(w *domain.Watchtable) (error) {
	e := errors.New("")
	return e
}

func (m *gcsqlRepository) DeleteAll(w *domain.Watchtable) (error) {
	e := errors.New("ff")
	return e
}


func (m *gcsqlRepository) GetAll() (domain.Watchtables ,error){
	return nil,nil
}
