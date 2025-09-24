package entity

type Chat struct {
	ID               int64  `db:"id" json:"id"`
	Type             string `db:"type" json:"type"`
	Title            string `db:"title" json:"title"`
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
	Username         string `db:"username" json:"username"`
	Bio              string `db:"bio" json:"bio,omitempty"`
	Photo            string `db:"photo" json:"photo,omitempty"`
	Description      string `db:"description" json:"description,omitempty"`
	InviteLink       string `db:"invite_link" json:"invite_link,omitempty"`
	PinnedMessage    string `db:"pinned_message" json:"pinned_message,omitempty"`
	Permissions      string `db:"permissions" json:"permissions,omitempty"`
	SlowMode         int64  `db:"slow_mode" json:"slow_mode_delay,omitempty"`
	StickerSet       string `db:"sticker_set" json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool   `db:"can_set_sticker_set" json:"can_set_sticker_set,omitempty"`
	LinkedChatID     int64  `db:"linked_chat_id" json:"linked_chat_id,omitempty"`
	ChatLocation     string `db:"chat_location" json:"location,omitempty"`
	Private          bool   `db:"private" json:"has_private_forwards,omitempty"`
	Protected        bool   `db:"protected" json:"has_protected_content,omitempty"`
	NoVoiceAndVideo  bool   `db:"no_voice_and_video" json:"has_restricted_voice_and_video_messages"`
	Created          string `db:"created" json:"created"`
}
type ChatList struct {
	Items []*Chat `json:"rows"`
}

type Chats interface {
	Insert(*Chat) error
}
