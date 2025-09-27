package middle

import (
	"errors"
	"net/mail"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
	"modernc.org/sqlite"
)

func insertSender(c tele.Context, app entity.Application) {
	if sender := c.Sender(); sender != nil {
		userBot := &entity.TelebotUser{
			ID:              sender.ID,
			FirstName:       sender.FirstName,
			LastName:        sender.LastName,
			Username:        sender.Username,
			LanguageCode:    sender.LanguageCode,
			IsBot:           sender.IsBot,
			IsPremium:       sender.IsPremium,
			AddedToMenu:     sender.AddedToMenu,
			CanJoinGroups:   sender.CanJoinGroups,
			CanReadMessages: sender.CanReadMessages,
			SupportsInline:  sender.SupportsInline,
		}
		if err := app.GetRepo().GetTelebotUsers().Insert(userBot); err != nil {
			e := errors.Unwrap(err)
			if er, ok := e.(*sqlite.Error); ok {
				if er.Code() != 1555 { // 1555 - code:1555 msg:constraint failed: UNIQUE constraint failed: chats.id (1555)
					app.GetLogger().Errorf("updates:insertSender %s", err.Error())
				}
			} else {
				app.GetLogger().Errorf("updates:insertSender %s", err.Error())
			}
		}
		app.GetLogger().Debugf("updates:insertSender [%v]%s", sender.ID, sender.Username)
	}
}

func insertChat(c tele.Context, app entity.Application) {
	if chat := c.Chat(); chat != nil {
		chatBot := &entity.Chat{
			ID:               chat.ID,
			Type:             string(chat.Type),
			Title:            chat.Title,
			FirstName:        chat.FirstName,
			LastName:         chat.LastName,
			Username:         chat.Username,
			Bio:              chat.Bio,
			Photo:            "",
			Description:      chat.Description,
			InviteLink:       chat.InviteLink,
			PinnedMessage:    "chat.PinnedMessage.Text",
			Permissions:      "chat.Permissions",
			SlowMode:         int64(chat.SlowMode),
			StickerSet:       chat.StickerSet,
			CanSetStickerSet: chat.CanSetStickerSet,
			LinkedChatID:     chat.LinkedChatID,
			ChatLocation:     "chat.ChatLocation.Address",
			Private:          chat.Private,
			Protected:        chat.Protected,
			NoVoiceAndVideo:  chat.NoVoiceAndVideo,
		}
		if err := app.GetRepo().GetChats().Insert(chatBot); err != nil {
			e := errors.Unwrap(err)
			if er, ok := e.(*sqlite.Error); ok {
				if er.Code() != 1555 { // 1555 - code:1555 msg:constraint failed: UNIQUE constraint failed: chats.id (1555)
					app.GetLogger().Errorf("updates:insertChat %s", err.Error())
				}
			} else {
				app.GetLogger().Errorf("updates:insertChat %s", err.Error())
			}
		}
	}
}

func checkSender(id int64, app entity.Application) bool {
	if user, err := app.GetRepo().GetTelebotUsers().GetById(id); err != nil {
		return false
	} else {
		return user != nil
	}
}

// проверяем состояния по ключам из таблицы не просто присутствие любого
func checkAllStates(um *entity.UserStatesMaps, app entity.Application) bool {
	defer app.GetRecovery().RecoverLog("middle:checkAllStates")
	if um == nil {
		return false
	}
	examens := app.GetRepo().GetKeyStates().GetExamenKeys()
	for _, ex := range examens {
		if val, ok := um.Examens[ex]; ok {
			if !checkStateDate(val, app) {
				return false
			}
		} else {
			return false
		}
	}
	intros := app.GetRepo().GetKeyStates().GetIntroKeys()
	for _, ex := range intros {
		if val, ok := um.Intro[ex]; ok {
			return checkIntro(val)
		} else {
			return false
		}
	}
	return true
}

// проверка что это дата и она в будущем
func checkStateDate(dt string, app entity.Application) bool {
	defer app.GetRecovery().RecoverLog("middle:checkStateDate")
	l := app.GetConfiguration().Layouts.TimeLayoutDay
	nowDate := time.Now().Local()
	if date, err := time.Parse(l, dt); err != nil {
		return false
	} else {
		days := int(date.Sub(nowDate).Hours() / 24)
		if days < 0 {
			return false
		}
	}
	return true
}

// проверка что это дата и она в будущем
func checkEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func checkIntro(val string) bool {
	return val != ""
}

// для режимов кроме редактора экзаменов пытаемся сбросить клавиатуру
// и удалить сообщение об этом
func clearKeyboard(c tele.Context) {
	menuClear := &tele.ReplyMarkup{RemoveKeyboard: true}
	if msg, err := c.Bot().Send(c.Sender(), "reset", menuClear, tele.ModeMarkdownV2); err == nil {
		c.Bot().Delete(msg)
	}
}
