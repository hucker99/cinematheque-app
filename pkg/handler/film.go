package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hucker99/cinematheque-app/model"
	"net/http"
	"strconv"
)

func (h *Handler) createFilm(c *gin.Context) {
	var input model.Film
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Film.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllFilms(c *gin.Context) {
	sortBy := c.Param("sort_by")
	if sortBy == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort_by param")
		return
	}

	films, err := h.services.Film.GetAll(sortBy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, films)
}

func (h *Handler) getFilmByFragment(c *gin.Context) {
}

func (h *Handler) updateFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateFilmInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Film.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) deleteFilm(c *gin.Context) {
	filmId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Film.Delete(filmId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": filmId,
	})
}
