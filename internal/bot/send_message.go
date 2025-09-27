package bot

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

// const (
//
//	ModeDefault    ParseMode = ""
//	ModeMarkdown   ParseMode = "Markdown"
//	ModeMarkdownV2 ParseMode = "MarkdownV2"
//	ModeHTML       ParseMode = "HTML"
//
// )
func (b Bot) SendMessageUser(uid int64, message string, mode string) error {
	user := &telebot.User{ID: uid}
	b.App.GetLogger().Infof("bot send message to %v mode %s", uid, mode)
	if _, err := b.Send(user, message, mode); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (b Bot) SendMessageAdmin(message string, mode string) error {
	user := &telebot.User{ID: b.App.GetConfiguration().AdminId}
	b.App.GetLogger().Infof("bot send message to admin %v mode %s", user.ID, mode)
	if _, err := b.Send(user, message, mode); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (b Bot) SendMessageGroup(message string, mode string) error {
	user := &telebot.User{ID: b.App.GetConfiguration().GroupId}
	b.App.GetLogger().Infof("bot send message to group %v mode %s", user.ID, mode)
	if _, err := b.Send(user, message, mode); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
func (b Bot) SendMessageChannel(message string, mode string) error {
	user := &telebot.User{ID: b.App.GetConfiguration().ChannelId}
	b.App.GetLogger().Infof("bot send message to channel %v mode %s", user.ID, mode)
	if _, err := b.Send(user, message, mode); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
