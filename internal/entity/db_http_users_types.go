package entity

type User struct {
	ClientId int64  `db:"client_id" json:"clientId"`
	Login    string `db:"login" json:"login"`
	Passwd   string `db:"passwd" json:"passwd"`
	Active   bool   `db:"active" json:"active"`
	IsAdmin  bool   `db:"is_admin" json:"isAdmin"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Rem      string `db:"rem" json:"rem"`
}

type UserList struct {
	Items []*User `json:"rows"`
}

type Users interface {
	Get() error
	Insert(*User) error
	Delete(*User) error
	Update(*User) error
	GetItems() []*User
	UpdateActive(*User) error

	GetByLogin(login string) (*User, error)
	GetByLoginAdmin(login string) (*User, error)
}
