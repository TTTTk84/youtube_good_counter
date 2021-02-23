package domain

import "net/http"

type Watchtable struct {
	Title string
	Url		string
	LikedAt	string
}

type Watchtables []Watchtable

type GCsqlRepository interface {
	Store(*Watchtable) error
	DeleteAll(Watchtables) error
	GetAll() (Watchtables, error)
}

type GCUsecase interface {
	AddGC(*Watchtable) (error)
	GetAll() (Watchtables, error)
	JsonParse(*http.Request) (*Watchtable, error)
}
