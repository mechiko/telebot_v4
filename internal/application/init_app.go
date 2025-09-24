package application

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/logrecover"
	"github.com/mechiko/telebot_v4/pkg/zaplog"
)

func initApplication() (*applicationType, error) {
	defer logrecover.RecoverFmt("initApplication()")

	if !entity.Supported {
		return nil, fmt.Errorf("initApplication not supported OS")
	}
	if entity.Linux {
		// linux create paths
		if _, err := os.Stat(entity.ConfigPath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			os.Mkdir(entity.ConfigPath, fs.FileMode(entity.PosixChownPath))
		}
		if _, err := os.Stat(entity.DbPath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			os.Mkdir(entity.DbPath, fs.FileMode(entity.PosixChownPath))
		}
		if _, err := os.Stat(entity.DbPath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			os.Mkdir(entity.LogPath, fs.FileMode(entity.PosixChownPath))
		}
		os.Chown(entity.ConfigPath, entity.PosixUserUIDGUID, entity.PosixUserUIDGUID)
		os.Chmod(entity.ConfigPath, fs.FileMode(entity.PosixChownPath))
		os.Chown(entity.DbPath, entity.PosixUserUIDGUID, entity.PosixUserUIDGUID)
		os.Chmod(entity.DbPath, fs.FileMode(entity.PosixChownPath))
		os.Chown(entity.LogPath, entity.PosixUserUIDGUID, entity.PosixUserUIDGUID)
		os.Chmod(entity.LogPath, fs.FileMode(entity.PosixChownPath))
	}

	applicationInstance := &applicationType{
		name:          "default",
		configuration: &entity.Configuration{},
	}
	applicationInstance.recovery = logrecover.NewRecoveryInterface(zaplog.Logger.Sugar())
	applicationInstance.logger = zaplog.Logger

	if err := applicationInstance.initConfig(); err != nil {
		return applicationInstance, fmt.Errorf("initConfig() %w", err)
	}
	applicationInstance.layout = applicationInstance.configuration.Layouts.TimeLayout

	if err := applicationInstance.initDb(entity.DbPath); err != nil {
		return applicationInstance, fmt.Errorf("initDb() %w", err)
	}
	// if err := appInst.initDb2(); err != nil {
	// 	return appInst, err
	// }
	// if err := applicationInstance.config.Set("application.console", isConsole(), true); err != nil {
	// 	zaplog.Logger.Errorf("app error %s", err.Error())
	// 	// applicationInstance.ErrorLog().AnErr("app error", err).Send()
	// }
	return applicationInstance, nil
}
