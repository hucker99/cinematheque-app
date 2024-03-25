package service

import (
	"github.com/hucker99/cinematheque-app/model"
	"github.com/hucker99/cinematheque-app/pkg/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) Create(film model.Film) (int, error) {
	return s.repo.Create(film)
}

func (s *FilmService) GetAll(sortBy string) ([]model.Films, error) {
	return s.repo.GetAll(sortBy)
}

func (s *FilmService) Update(id int, input model.UpdateFilmInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}

func (s *FilmService) Delete(filmId int) error {
	return s.repo.Delete(filmId)
}
