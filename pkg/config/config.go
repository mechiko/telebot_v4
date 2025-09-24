package config

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/pkg/zaplog"
	"github.com/spf13/viper"
)

type Config struct {
	entity.ConfigInterface
	app            entity.Application
	viper          *viper.Viper
	configuration  interface{}
	configFileName string
}

// var TomlConfig []byte
var YamlConfig []byte

const defaultConfigName = "config"

// создание нового экземпляра конфига реентерабельно возвращается интерфейс
// dbg ...string для отладки передаем строки первая печатается в консоль.. это видимо был тест параметров, пусть так пока
func NewInstance(app entity.Application, cfgName string, configuration interface{}, dbg ...string) (*Config, error) {
	defer RecoverFmt("config:NewInstance")
	var cfg *Config
	var err error
	if len(dbg) > 0 {
		zaplog.Logger.Sugar().Infof("package:Config file:configuration.go func:GetInstance() msg:doOnce invoked by %v", dbg[0])
	}
	if cfg, err = initConfiguration(app, cfgName, configuration); err != nil {
		return nil, fmt.Errorf("config:NewInstance() %w", err)
	}
	if len(dbg) > 0 {
		app.GetLogger().Infof("Dbg=%v\n", dbg[0])
		pc, file, no, ok := runtime.Caller(1)
		if ok {
			details := runtime.FuncForPC(pc)
			zaplog.Logger.Sugar().Infof("package:Config file:configuration.go func:GetInstance() msg:called from %s#%d", file, no)
			if details != nil {
				zaplog.Logger.Sugar().Infof("package:Config file:configuration.go func:GetInstance() msg:called from %s", details.Name())
			}
		}
	}
	return cfg, nil
}

func initConfiguration(app entity.Application, cfgName string, configuration interface{}) (*Config, error) {
	defer RecoverFmt("initConfiguration")

	configName := defaultConfigName
	if cfgName != "" {
		configName = cfgName
	}
	configName = path.Join(entity.ConfigPath, configName)

	viperOrigin := viper.GetViper()
	configFileName := configName + ".yaml"

	if entity.Windows {
		viperOrigin.SetConfigName(configName)
		viperOrigin.SetConfigType("yaml")
		viperOrigin.AddConfigPath(".")
	} else {
		viperOrigin.SetConfigType("yaml")
		viperOrigin.SetConfigFile(configFileName)
		viperOrigin.SetConfigPermissions(0644)
		viperOrigin.AddConfigPath(entity.ConfigPath)
	}
	if err := viperOrigin.MergeConfig(strings.NewReader(string(YamlConfig))); err != nil {
		return nil, fmt.Errorf("viperOrigin.MergeConfig() %w", err)
	}

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			zaplog.Logger.Sugar().Errorf("Config file ('%s') not found", configFileName)
		} else {
			zaplog.Logger.Sugar().Errorf("viper.MergeInConfig type %T err  %+v ", err, err)
		}
	}

	viperOrigin.AutomaticEnv()
	viperOrigin.SetEnvPrefix("BOT")
	viperOrigin.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(configuration); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal(configuration) %w", err)
	}

	cfg := &Config{
		app:            app,
		configuration:  configuration,
		configFileName: configFileName,
		viper:          viperOrigin,
	}
	viperOrigin.SafeWriteConfig()

	return cfg, nil
}
