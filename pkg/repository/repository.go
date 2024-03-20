package repository

type Authorization interface {
}

type Film interface {
}

type Actor interface {
}

type Repository struct {
	Authorization
	Film
	Actor
}

func NewRepository() *Repository {
	return &Repository{}
}
