package db

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/telebot_v4/internal/entity"
)

type DbRoutes struct {
	app     entity.Application
	useCase entity.UseCase
}

func New(handler *echo.Group, app entity.Application) {
	// r := &DbRoutes{
	// 	app:     app,
	// 	useCase: usecase.New(app),
	// }

	h := handler.Group("/db")
	{
		h.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
	}
}
