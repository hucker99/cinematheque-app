package model

import "time"

type Actor struct {
	Id           uint64 `json:"id,omitempty"`
	Name         string `json:"name"`
	Gender       string `json:"gender" valid:"in(Male|Female)"`
	BirthdayStr  string `json:"birthday"`
	BirthdayTime time.Time
}

type Film struct {
	Id              uint64 `json:"id,omitempty"`
	Name            string `json:"name"`
	ReleaseDateStr  string `json:"release_date"`
	ReleaseDateTime time.Time
	Rating          int   `json:"rating"`
	Actors          []int `json:"actors"`
}

type User struct {
	ID       int    `json:"-"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password"`
	Role     string `json:"-"`
}

type Session struct {
	UID        int       `json:"-"`
	Cookie     string    `json:"cookie"`
	ExpireDate time.Time `json:"expire_date"`
}
