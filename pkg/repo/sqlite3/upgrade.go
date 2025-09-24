package sqlite3

import (
	// _ "embed"
	"fmt"
	// этот драйвер не зависит от CGO поэтому не проблема для 64 бит
	// _ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

// upgradeDb() обновляем БД и номер версии туда скрипт s.UpdateDb
func (r *repository) upgradeDb() error {
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.App.GetDbService().Close()
	}()
	// сначала пытаемся обновить БД
	if _, err = db.Exec(UpdateDB); err != nil {
		return fmt.Errorf("db.Exec(UpdateDB) %w", err)
	}
	// теперь обновляем версию
	if _, err = db.Exec(setVersion, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(setVersion, ExeVersion) %w", err)
	}
	return err
}
