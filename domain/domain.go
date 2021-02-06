package domain

type Watchtable struct {
	Title string
	Url		string
	LikedAt	string
}

type Watchtables []Watchtable

type GCsqlRepository interface {
	Store(*Watchtable) error
	DeleteAll(*Watchtable) error
	GetAll() (Watchtables, error)
}

type GCUsecase interface {
	AddGC()
	GCOnceday()
	GCOnceWeek()
}
