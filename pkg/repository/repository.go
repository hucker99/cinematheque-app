package repository

import (
	"github.com/hucker99/cinematheque-app/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
}

type Film interface {
	Create(film model.Film) (int, error)
	GetAll(sortBy string) ([]model.Films, error)
	Update(id int, input model.UpdateFilmInput) error
	Delete(filmId int) error
}

type Actor interface {
	Create(actor model.Actor) (int, error)
	GetAll() ([]model.ActorsFilmsRelations, error)
	Update(id int, input model.UpdateActorInput) error
	Delete(actorId int) error
}

type Repository struct {
	Authorization
	Film
	Actor
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Film:          NewFilmPostgres(db),
		Actor:         NewActorPostgres(db),
	}
}
