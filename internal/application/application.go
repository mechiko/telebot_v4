package application

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/logrecover"
	"go.uber.org/zap"
)

type applicationType struct {
	entity.Application
	name          string
	config        entity.ConfigInterface
	logger        *zap.Logger
	configuration *entity.Configuration
	db            entity.DbService
	repo          entity.Repo
	recovery      entity.RecoverInterface
	layout        string
	bot           entity.Bot
	// Settings
	pwd string
}

var _ entity.Application = &applicationType{}

func NewApplication() (entity.Application, error) {
	defer logrecover.RecoverFmt("NewApplication()")
	if applicationInstance, err := initApplication(); err != nil {
		return nil, fmt.Errorf("application:NewApplication() %w", err)
	} else {
		return applicationInstance, nil
	}
}
