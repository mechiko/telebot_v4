package entity

// таблица для хранения постоянных значений на приложение
// меняющихся по надобности, пока таких нет
type AppState struct {
	Module string `db:"module"`
	Key    string `db:"name"`
	Value  string `db:"value"`
}

type AppStateList struct {
	Items []AppState
}

type AppStates interface {
	Set(module string, key string, value string) error
	Get(module string, key string) (string, error)
}
