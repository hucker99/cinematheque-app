package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hucker99/cinematheque-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		films := api.Group("/films")
		{
			films.POST("/", h.createFilm)
			films.GET("/", h.getAllFilms)
			films.GET("/:id", h.getFilmByFragment)
			films.PUT("/:id", h.updateFilm)
			films.DELETE("/:id", h.deleteFilm)

			actor := films.Group(":id/actors")
			{
				actor.POST("/", h.createActor)
				actor.GET("/", h.getAllActors)
				actor.PUT("/:actor_id", h.updateActor)
				actor.DELETE("/:actor_id", h.deleteActor)
			}
		}
	}
	return router
}
