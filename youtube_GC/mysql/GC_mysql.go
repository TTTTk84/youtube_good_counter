package mysql

import (
	"database/sql"
	"errors"

	"github.com/TTTTk84/youtube_good_counter/domain"
)
type gcsqlRepository struct {
	Conn *sql.DB
}

const (
	queryInsertWT = "INSERT INTO watchtables (title, url, likedAt) VALUES(?,?,?)"
	queryDeleteAllWT = "TRUNCATE TABLE watchtables"
	queryGetAllWT= "SELECT * FROM watchtables"
)

func NewGC_mysql(Conn *sql.DB) domain.GCsqlRepository {
	return &gcsqlRepository{Conn}
}

func (m *gcsqlRepository) Store(w *domain.Watchtable) (error) {
	stmt, err := m.Conn.Prepare(queryInsertWT)
	if err != nil{
		return errors.New("err: failed insert into")
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		w.Title,
		w.Url,
		w.LikedAt,
	)
	if err != nil{
		return errors.New("err: failed insert into")
	}

	return nil
}

func (m *gcsqlRepository) DeleteAll(w *domain.Watchtable) (error) {
	stmt, err := m.Conn.Prepare(queryDeleteAllWT)
	if err != nil {
		return errors.New("err: failed delete all")
	}
	defer stmt.Close()

	return nil
}


func (m *gcsqlRepository) GetAll() (domain.Watchtables ,error){
	stmt, err := m.Conn.Prepare(queryGetAllWT)
	if err != nil {
		return nil, errors.New("err: failed get all")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("err: failed get rows")
	}


	var wts domain.Watchtables
	for rows.Next() {
		var wt domain.Watchtable

		err := rows.Scan(
			&wt.Title,
			&wt.Url,
			&wt.LikedAt,
		)
		if err != nil{
			return nil, errors.New("err: failed row scan")
		}
		wts = append(wts, wt)
	}

	return wts, nil
}
