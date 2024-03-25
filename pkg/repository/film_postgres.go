package repository

import (
	"fmt"
	"github.com/hucker99/cinematheque-app/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type FilmPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *FilmPostgres {
	return &FilmPostgres{db: db}
}

func (r *FilmPostgres) Create(film model.Film) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var filmID int
	createFilmQuery := fmt.Sprintf("INSERT INTO %s (title, release_date, rating) VALUES ($1, $2, $3) RETURNING id", filmTable)
	row := tx.QueryRow(createFilmQuery, film.Title, film.ReleaseDateStr, film.Rating)
	if err = row.Scan(&filmID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createFilmsActorsQuery := fmt.Sprintf("INSERT INTO %s (film_id, actor_id) VALUES ($1, $2)", filmsActorsTable)
	for _, actorID := range film.Actors {
		_, err = tx.Exec(createFilmsActorsQuery, filmID, actorID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return filmID, tx.Commit()
}

func (r *FilmPostgres) GetAll(sortBy string) ([]model.Films, error) {
	var films []model.Films

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY CASE WHEN $1 = 'title' THEN title WHEN $1 = 'release_date' THEN release_date ELSE rating END DESC",
		filmTable)
	if err := r.db.Select(&films, query, sortBy); err != nil {
		return nil, err
	}

	return films, nil
}

func (r *FilmPostgres) Update(id int, input model.UpdateFilmInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.ReleaseDate != nil {
		setValues = append(setValues, fmt.Sprintf("release_date=$%d", argId))
		args = append(args, *input.ReleaseDate)
		argId++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	updateFilmQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		filmTable, setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateFilmQuery: %s", updateFilmQuery)
	logrus.Debugf("args: %s", args)
	_, err = tx.Exec(updateFilmQuery, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	// updating films_actors table
	updateFilmsActorsQuery := fmt.Sprintf("UPDATE %s SET actor_id = $1 WHERE film_id = $2",
		filmsActorsTable)
	logrus.Debugf("updateFilmsActorsQuery: %s", updateFilmsActorsQuery)
	for _, actorID := range *input.Actors {
		_, err = tx.Exec(updateFilmsActorsQuery, actorID, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *FilmPostgres) Delete(filmId int) error {
	//query := fmt.Sprintf("DELETE FROM %s ft USING %s fat WHERE ft.id = fat.film_id AND ft.id=$1 AND fat.film_id=$1", filmTable, filmsActorsTable)
	//_, err := r.db.Exec(query, filmId)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteFilmQuery := fmt.Sprintf("DELETE FROM %s at WHERE at.id = $1", filmTable)
	_, err = tx.Exec(deleteFilmQuery, filmId)
	if err != nil {
		return err
	}

	deleteFilmsActorsQuery := fmt.Sprintf("DELETE FROM %s fat WHERE fat.film_id = $1", filmsActorsTable)
	_, err = tx.Exec(deleteFilmsActorsQuery, filmId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
