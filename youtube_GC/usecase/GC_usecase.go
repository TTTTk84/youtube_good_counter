package usecase

import "github.com/TTTTk84/youtube_good_counter/domain"

type gcUsecase struct {
	gcsqlrepo domain.GCsqlRepository
}

func NewGCUsecase(gu domain.GCsqlRepository) domain.GCUsecase {
	return &gcUsecase{gu}
}

func (g *gcUsecase) AddGC() {

}

func (g *gcUsecase) GCOnceday() {

}

func (g *gcUsecase) GCOnceWeek() {

}
