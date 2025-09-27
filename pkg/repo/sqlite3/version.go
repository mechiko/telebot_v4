package sqlite3

import (
	_ "embed"
	"fmt"

	"github.com/hashicorp/go-version"
	// этот драйвер не зависит от CGO поэтому не проблема для 64 бит
	// _ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

//goland:noinspection ALL
const createVersion string = `
CREATE TABLE if not exists dboptions (
	name TEXT NOT NULL DEFAULT '',
	value TEXT NOT NULL DEFAULT '',
	PRIMARY KEY (name)
);`

const setVersion string = `INSERT OR REPLACE INTO dboptions (name, value) VALUES ('version',?)`

const getVersion string = `select value from dboptions where name = 'version' limit 1;`

const ExeVersion string = "0.0.4"

//go:embed createDB.sql
var CreateDB string

//go:embed updateDB.sql
var UpdateDB string

//go:embed index.sql
var IndexDB string

//go:embed insert.sql
var InsertDB string

// true если версия новее
func (r *repository) CheckVersionDb() (bool, error) {
	var ver string
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return false, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.App.GetDbService().Close()
	}()
	// проверяем есть ли вообще версия в БД
	if err := db.Get(&ver, getVersion); err != nil {
		return false, fmt.Errorf("db.Get() %w", err)
	}
	v1, _ := version.NewVersion(ExeVersion)
	v2, _ := version.NewVersion(ver)
	if v2.LessThan(v1) {
		return true, nil
	}
	return false, err
}
