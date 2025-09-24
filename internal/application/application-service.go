package application

import (
	"fmt"
	"os"

	"github.com/mechiko/telebot_v4/internal/logrecover"
)

const DateFormat = "2006.01.02"

// это для повторной инициализации конфигурации повторное чтение данных
func (a *applicationType) InitConfiguration() error {
	defer logrecover.RecoverFmt("(a *application) InitConfiguration()")
	var err error
	var cfg = a.configuration

	if a.pwd, err = os.Getwd(); err != nil {
		return fmt.Errorf("InitConfiguration() %w", err)
	}
	if err = a.GetConfig().Unmarshal(cfg); err != nil {
		return fmt.Errorf("InitConfiguration() %w", err)
	}
	return nil
}

func (a *applicationType) SaveConfig() error {
	if err := a.config.SaveAs("cccc.ccc"); err != nil {
		return fmt.Errorf("application:SaveConfig() %w", err)
	}
	return nil
}

func (a *applicationType) GetPwd() string {
	return a.pwd
}

func (a *applicationType) GetBaseUrl() string {
	// var cfg = &entity.Configuration{}
	// config.GetConfig().Viper.Unmarshal(cfg)

	uri := "https://" + a.configuration.Hostname
	return uri
}
