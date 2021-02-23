package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TTTTk84/youtube_good_counter/domain"
)

type gcUsecase struct {
	gcsqlrepo domain.GCsqlRepository
}

func NewGC_Usecase(gu domain.GCsqlRepository) domain.GCUsecase {
	return &gcUsecase{gu}
}

func (g *gcUsecase) AddGC(w *domain.Watchtable) (error){
	err := g.gcsqlrepo.Store(w)
	if err != nil {
		return err
	}

	return nil
}

func (g *gcUsecase) GetAll() (domain.Watchtables, error){
	var wts domain.Watchtables
	wts, err :=  g.gcsqlrepo.GetAll()
	if err != nil {
		return wts, err
	}

	err = g.gcsqlrepo.DeleteAll(wts)
	if err != nil {
		return wts, err
	}

	return wts, nil
}

func (g *gcUsecase) JsonParse(r *http.Request) (*domain.Watchtable, error) {
	wt := &domain.Watchtable{}
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return wt,err
	}

	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	err = json.Unmarshal(body[:length], &wt)
	if err != nil {
		return wt,err
	}

	return wt, nil
}
