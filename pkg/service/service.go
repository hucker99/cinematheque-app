package service

import (
	"github.com/hucker99/cinematheque-app/model"
	"github.com/hucker99/cinematheque-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Film
	Actor
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewAuthService(repos.Authorization),
		NewFilmService(repos.Film),
		NewActorService(repos.Actor),
	}
}
