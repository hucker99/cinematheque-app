package service

import (
	"github.com/hucker99/cinematheque-app/model"
	"github.com/hucker99/cinematheque-app/pkg/repository"
)

type ActorService struct {
	repo repository.Actor
}

func NewActorService(repo repository.Actor) *ActorService {
	return &ActorService{repo: repo}
}

func (s *ActorService) Create(actor model.Actor) (int, error) {
	return s.repo.Create(actor)
}

func (s *ActorService) GetAll() ([]model.ActorsFilmsRelations, error) {
	return s.repo.GetAll()
}

func (s *ActorService) Update(id int, input model.UpdateActorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}

func (s *ActorService) Delete(actorId int) error {
	return s.repo.Delete(actorId)
}
