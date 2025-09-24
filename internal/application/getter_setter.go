package application

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"go.uber.org/zap"
)

func (a *applicationType) GetConfig() entity.ConfigInterface {
	return a.config
}

func (a *applicationType) GetConfiguration() *entity.Configuration {
	return a.configuration
}

func (a *applicationType) GetDbService() entity.DbService {
	// тут бы ошибку может какую а так это пустой код
	// if a.Db == nil {
	// 	return nil
	// }
	return a.db
}

func (a *applicationType) SetRepo(repo entity.Repo) {
	a.repo = repo
}

func (a *applicationType) GetRepo() entity.Repo {
	return a.repo
}

func (a *applicationType) InitDb() error {
	p := entity.DbPath
	return a.initDb(p)
}

func (a *applicationType) GetRecovery() entity.RecoverInterface {
	return a.recovery
}

func (a *applicationType) Shutdown() {
	// a.logger.Dispose()
	entity.AppInterrupt <- 0
}

func (a *applicationType) Restart() {
	// a.gui.Shutdown()
	// a.logger.Dispose()
	entity.AppInterrupt <- 1
}

func (a *applicationType) RestartConsole() {
	// a.logger.Dispose()
	entity.AppInterrupt <- 1
}

func (a *applicationType) GetLogger() *zap.SugaredLogger {
	return a.logger.Sugar()
}

func (a *applicationType) SetBot(b entity.Bot) {
	a.bot = b
}

func (a *applicationType) GetBot() entity.Bot {
	return a.bot
}
