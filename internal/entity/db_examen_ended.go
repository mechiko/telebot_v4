package entity

type ExamenEnded struct {
	UserId  int64  `db:"user_id" json:"userId"`
	Examen  string `db:"key" json:"examen"`
	Date    string `db:"value" json:"date"`
	Created string `db:"created" json:"created"`
}

type ExamenEndedList struct {
	Items []*ExamenEnded
}

type ExamenEndeds interface {
	Get() (*ExamenEndedList, error)
	Insert(*ExamenEnded) error
}
