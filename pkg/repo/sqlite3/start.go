package sqlite3

import (
	"fmt"
)

func (r *repository) Start() error {
	defer r.App.GetRecovery().RecoverLog("Start()")
	dbService := r.App.GetDbService()
	// создаем структуру БД если файл БД был только что создан
	if dbService.IsCreated() {
		if err := r.create(); err != nil {
			return fmt.Errorf("r.create() %w", err)
		}
		if err := r.InsertUsers(); err != nil {
			return fmt.Errorf("InsertUsers() %w", err)
		}
		if err := r.InsertStateKeys(); err != nil {
			return fmt.Errorf("%w", err)
		}
		dbService.SetCreated()
	} else {
		// БД при запуске присутствовала
		// проверяем версию и обновляем если надо
		needUpgrade, err := r.CheckVersionDb()
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		if needUpgrade {
			if err := r.upgradeDb(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}
	// каждый раз обновляем ключи для состояний
	// при смене этих ключей надо будет как то править данные но это отдельная история
	// if err := r.InsertStateKeys(); err != nil {
	// 	return fmt.Errorf("%w", err)
	// }
	// пересоздаем при старте View
	r.App.GetLogger().Debugf("!create views in db!")
	if err := r.GetViews().Create(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if r.App.GetConfiguration().Debug {
		r.CheckOnStart()
	}
	return nil
}
