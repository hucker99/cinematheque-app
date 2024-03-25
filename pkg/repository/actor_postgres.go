package repository

import (
	"fmt"
	"github.com/hucker99/cinematheque-app/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type ActorPostgres struct {
	db *sqlx.DB
}

func NewActorPostgres(db *sqlx.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (r *ActorPostgres) Create(actor model.Actor) (int, error) {
	var actorID int
	createActorQuery := fmt.Sprintf("INSERT INTO %s (name, gender, birthday) VALUES ($1, $2, $3) RETURNING id", actorTable)
	row := r.db.QueryRow(createActorQuery, actor.Name, actor.Gender, actor.BirthdayStr)
	if err := row.Scan(&actorID); err != nil {
		return 0, err
	}
	return actorID, nil
}

func (r *ActorPostgres) GetAll() ([]model.ActorsFilmsRelations, error) {
	var actorsFilms []model.ActorsFilmsRelations

	query := fmt.Sprintf("SELECT at.name, ft.title FROM %s at JOIN %s fat ON at.id = fat.actor_id JOIN %s ft ON fat.film_id = ft.id",
		actorTable, filmsActorsTable, filmTable)
	if err := r.db.Select(&actorsFilms, query); err != nil {
		return nil, err
	}

	return actorsFilms, nil
}

func (r *ActorPostgres) Update(id int, input model.UpdateActorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Gender != nil {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *input.Gender)
		argId++
	}

	if input.Birthday != nil {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, *input.Birthday)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s at SET %s WHERE at.id = $%d",
		actorTable, setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ActorPostgres) Delete(actorId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteActorQuery := fmt.Sprintf("DELETE FROM %s at WHERE at.id = $1", actorTable)
	_, err = tx.Exec(deleteActorQuery, actorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteFilmsActorQuery := fmt.Sprintf("DELETE FROM %s fat WHERE fat.actor_id = $1", filmsActorsTable)
	_, err = tx.Exec(deleteFilmsActorQuery, actorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
