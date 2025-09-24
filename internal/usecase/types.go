// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type usecase struct {
	App entity.Application
}

// asserts
var _ entity.UseCase = (*usecase)(nil)

// New -.
func New(a entity.Application) entity.UseCase {
	return &usecase{
		App: a,
	}
}
