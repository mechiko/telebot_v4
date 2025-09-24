package entity

type TelebotUser struct {
	ID              int64  `db:"id" json:"id"`
	FirstName       string `db:"first_name" json:"firstName"`
	LastName        string `db:"last_name" json:"lastName"`
	Username        string `db:"username" json:"userName"`
	LanguageCode    string `db:"language_code" json:"languageCode"`
	IsBot           bool   `db:"is_bot" json:"isBot"`
	IsPremium       bool   `db:"is_premium" json:"isPremium"`
	AddedToMenu     bool   `db:"added_to_menu" json:"addedToMenu"`
	CanJoinGroups   bool   `db:"can_join_groups" json:"canJoinGroups"`
	CanReadMessages bool   `db:"can_read_messages" json:"canReadMessages"`
	SupportsInline  bool   `db:"supports_inline" json:"supportsInline"`
	Created         string `db:"created" json:"created"`
	IsAdmin         bool   `db:"is_admin" json:"isAdmin"`
	Masters         string `db:"masters" json:"masters"`
	Ident           string `db:"ident" json:"ident"`
}
type TelebotUserList struct {
	Items []*TelebotUser `json:"rows"`
}

type TelebotUsers interface {
	GetList() (*TelebotUserList, error)
	Insert(*TelebotUser) error
	Update(u *TelebotUser) error
	Delete(u *TelebotUser) error
	GetById(id int64) (*TelebotUser, error)
}
