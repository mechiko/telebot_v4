package keystates

import (
	"fmt"

	"github.com/mechiko/telebot_v4/pkg/zaplog"
)

func (r *keystates) GetIntroKeys() []string {
	defer r.Repo.GetApplication().GetRecovery().RecoverLog("GetIntroKeys")
	ar := []string{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		zaplog.Logger.Sugar().Errorf("repo:keystates %s", err.Error())
		return ar
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select key from state_key where is_intro = 1;`

	err = db.Select(&ar, query)
	if err != nil {
		r.Repo.GetApplication().GetLogger().Errorf("repo:keystates %s", err.Error())
		return ar
	}

	return ar
}

func (r *keystates) GetExamenKeys() []string {
	defer r.Repo.GetApplication().GetRecovery().RecoverLog("GetExamenKeys")
	ar := []string{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		zaplog.Logger.Sugar().Errorf("repo:keystates %s", err.Error())
		return ar
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select key from state_key where is_examen = 1;`

	err = db.Select(&ar, query)
	if err != nil {
		r.Repo.GetApplication().GetLogger().Errorf("repo:keystates %s", err.Error())
		return ar
	}
	return ar
}

func (r *keystates) Get(userid int64, key string) (string, error) {
	value := ""
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return value, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select value from user_states where user_id = ? and key = ?;`

	err = db.Get(&value, query, userid, key)
	if err != nil {
		return value, fmt.Errorf("db.Select() %w", err)
	}
	return value, nil
}
