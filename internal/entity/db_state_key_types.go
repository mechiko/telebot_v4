package entity

type KeyState struct {
	Key         string `db:"key"`
	Description string `db:"description"`
	IsExamen    bool   `db:"is_examen"`
	IsIntro     bool   `db:"is_intro"`
}

type KeyStateList struct {
	Items []KeyState
}

type KeyStates interface {
	// GetAllKeys() []string
	GetIntroKeys() []string
	GetExamenKeys() []string
}
