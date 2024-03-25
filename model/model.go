package model

import (
	"errors"
	"time"
)

type Actor struct {
	Id          uint64 `json:"-" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Gender      string `json:"gender" db:"gender" binding:"required" valid:"in(male|female)"`
	BirthdayStr string `json:"birthday" db:"birthday" binding:"required" `
	// TODO
	BirthdayTime time.Time `json:"-" db:"-"`
}

type Film struct {
	Id             uint64 `json:"-" db:"id"`
	Title          string `json:"title" db:"title" binding:"required"`
	ReleaseDateStr string `json:"release_date" db:"release_date" binding:"required"`
	// TODO
	ReleaseDateTime time.Time `json:"-" db:"-"`
	Rating          int       `json:"rating" db:"rating" binding:"required"`
	Actors          []int     `json:"actors"`
}

type User struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required" valid:"email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"-"`
}

type Films struct {
	Id          uint64 `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Rating      int    `db:"rating"`
}

type ActorsFilmsRelations struct {
	ActorName string `db:"name"`
	FilmTitle string `db:"title"`
}

type UpdateActorInput struct {
	Name     *string `json:"name"`
	Gender   *string `json:"gender" valid:"in(male|female)"`
	Birthday *string `json:"birthday"`
}

func (a UpdateActorInput) Validate() error {
	if a.Name == nil && a.Gender == nil && a.Birthday == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateFilmInput struct {
	Title       *string `json:"title"`
	ReleaseDate *string `json:"release_date"`
	Rating      *int    `json:"rating"`
	Actors      *[]int  `json:"actors"`
}

func (a UpdateFilmInput) Validate() error {
	if a.Title == nil && a.ReleaseDate == nil && a.Rating == nil && a.Actors == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
