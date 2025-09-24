package entity

import "go.uber.org/zap"

type Application interface {
	GetConfig() ConfigInterface
	GetConfiguration() *Configuration
	GetDbService() DbService
	GetPwd() string
	SaveConfig() error
	GetBaseUrl() string
	DumpSql(s string)
	DumpSqlAppend(s string)
	DumpSqlClear()
	SetRepo(Repo)
	GetRepo() Repo
	InitDb() error
	GetRecovery() RecoverInterface
	Shutdown()
	Restart()
	RestartConsole()
	GetLogger() *zap.SugaredLogger
	SetBot(Bot)
	GetBot() Bot
}

var AppInterrupt = make(chan int, 2)
