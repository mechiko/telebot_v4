package chats

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *chats) Insert(h *entity.Chat) error {
	query := `INSERT INTO chats (id, type, title, first_name, last_name, username, bio, photo, description, invite_link, pinned_message, 
	permissions, slow_mode, sticker_set, can_set_sticker_set, linked_chat_id, chat_location, private, protected, no_voice_and_video)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("chats:insert %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, h.ID, h.Type, h.Title, h.FirstName, h.LastName, h.Username, h.Bio, h.Photo, h.Description, h.InviteLink, h.PinnedMessage,
		h.Permissions, h.SlowMode, h.StickerSet, h.CanSetStickerSet, h.LinkedChatID, h.ChatLocation, h.Private, h.Protected, h.NoVoiceAndVideo); err != nil {
		return fmt.Errorf("chats:insert db.Exec() %w", err)
	}
	return nil
}
