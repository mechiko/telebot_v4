package entity

type MasterUser struct {
	UserId            int64  `db:"id" json:"userId"`
	UserIdent         string `db:"ident" json:"userIdent"`
	IsAdmin           bool   `db:"is_admin" json:"isAdmin"`
	IdsSubordinate    string `db:"ids_subordinate" json:"idsSubordinate"`
	IdentsSubordinate string `db:"idents_subordinate" json:"identsSubordinate"`
}

type MasterUserList struct {
	Items []*MasterUser `json:"rows"`
}

type MasterUsers interface {
	GetAll() (*MasterUserList, error)
}
