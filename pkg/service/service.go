package service

import "github.com/hucker99/cinematheque-app/pkg/repository"

type Authorization interface {
}

type Film interface {
}

type Actor interface {
}

type Service struct {
	Authorization
	Film
	Actor
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
