package application

import (
	"fmt"
	"os"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/pkg/config"
)

func (a *applicationType) initConfig() error {
	// Configuration
	config.YamlConfig = entity.YamlConfig
	// tt, err := config.GetInstance("config", a.Configuration, "debug")
	tt, err := config.NewInstance(a, "config", a.configuration)
	if err != nil {
		fmt.Printf("app initConfig() error = %v\n", err.Error())
		os.Exit(1)
	}
	a.config = tt
	return nil
}
