package usecase

import "github.com/TTTTk84/youtube_good_counter/domain"

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

func (g *gcUsecase) GCOnceday() (domain.Watchtables, error){
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

func (g *gcUsecase) GCOnceWeek() {

}
