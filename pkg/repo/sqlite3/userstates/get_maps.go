package userstates

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// всегда возвращает список состояний всех пользователей
func (r *userstates) GetUserStateList() (*entity.UserStateList, error) {
	l := &entity.UserStateList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select u.user_id, u.key, u.value, k.is_examen, k.is_intro	from user_states u join state_key k on k.key = u.key;`

	err = db.Select(&l.Items, query)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return l, nil
}

// возвращает список состояний пользователя по ид
func (r *userstates) GetUserStateListByUser(userid int64) (*entity.UserStateList, error) {
	l := &entity.UserStateList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select u.key, u.value, k.is_examen, k.is_intro from user_states u join state_key k on k.key = u.key where u.user_id = ?; `

	err = db.Select(&l.Items, query, userid)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return l, nil
}
