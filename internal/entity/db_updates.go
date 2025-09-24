package entity

type Update struct {
	ID        int64  `db:"id" json:"id"`
	Message   string `db:"message" json:"message"`
	Sender    int64  `db:"sender_id" json:"sender_id"`
	Chat      int64  `db:"chat_id" json:"chat_id"`
	Recepient string `db:"recepient" json:"recepient_id"`
	Update    string `db:"update" json:"update"`
	Created   string `db:"created" json:"created"`
}
type UpdateList struct {
	Items []*Update `json:"rows"`
}

type Updates interface {
	Insert(*Update) error
}
