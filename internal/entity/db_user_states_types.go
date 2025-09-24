package entity

type UserState struct {
	UserId      int64  `db:"user_id"`
	Key         string `db:"key"`
	Value       string `db:"value"`
	Description string `db:"description"`
	IsExamen    bool   `db:"is_examen"`
	IsIntro     bool   `db:"is_intro"`
}

type UserStateList struct {
	Items []UserState
}

type UserStates interface {
	Get(userid int64, key string) (string, error)
	Set(userid int64, key string, value string) error
	GetUserStateList() (*UserStateList, error)
	GetUserStateListByUser(userid int64) (*UserStateList, error)
}
