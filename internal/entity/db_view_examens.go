package entity

type ExamenUser struct {
	ExamenKey      string `db:"examen" json:"examenKey"`
	UserName       string `db:"name" json:"userName"`
	UserId         int64  `db:"id" json:"userId"`
	UserIdent      string `db:"ident" json:"userIdent"`
	DateOrdered    string `db:"date_ordered" json:"dateOrdered"`
	BeforeDateDays int64  `db:"before_date_days" json:"beforeDateDays"`
	Date           string `db:"date" json:"date"`
	Masters        string `db:"masters_list" json:"masters"`
}

type ExamenUserList struct {
	Items []*ExamenUser `json:"rows"`
}

type ExamenUsers interface {
	GetAll() (*ExamenUserList, error)
	// GetAllEnded() (*ExamenUserList, error)
	// GetByMasterActive() (*ExamenUserList, error)
	// GetByUserActive() (*ExamenUserList, error)
	// GetByMaster() (*ExamenUserList, error)
	// GetByUser() (*ExamenUserList, error)
}
