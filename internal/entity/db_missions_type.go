package entity

type Mission struct {
	ID          int64  `db:"id" json:"id"`
	RecepientId int64  `db:"recepient_id" json:"recepient_id"`
	Start       string `db:"start" json:"start"`
	End         string `db:"end" json:"end"`
	Place       string `db:"place" json:"place"`
	Department  string `db:"departament" json:"departament"`
	Rem         string `db:"rem" json:"rem"`
	Target      string `db:"target" json:"target"`
	Created     string `db:"created" json:"created"`
	Reported    bool   `db:"reported" json:"reported"`
	Active      bool   `db:"active" json:"active"`
	ErrorTxt    string // это поле для бота
}

type MissionList struct {
	Items []*Mission `json:"rows"`
}

type Missions interface {
	Get() (*MissionList, error)
	GetById(id int64) (*Mission, error)
	GetActiveByUid(uid int64) (*Mission, error)
	Insert(*Mission) error
	Delete(ci *Mission) error
	DeleteById(id int64) error
	Update(m *Mission) error
	UpdateClearActive(m *Mission) error
	UpdateClearActiveById(int64) error
}
