package entity

import "strconv"

type DialogState interface {
}

type StoredMessage struct {
	MessageID int64
	ChatID    int64
}

func (x *StoredMessage) MessageSig() (string, int64) {
	return strconv.FormatInt(x.MessageID, 10), x.ChatID
}

type StateDialog struct {
	// UserId int64
	User      *TelebotUser
	ChatId    int64
	Text      string
	ErrorText string
	// Menu    *tele.ReplyMarkup
	Mode                 DialogMode
	Stage                ModeStage
	Mission              *Mission
	MissionDialogMsgs    []*StoredMessage // dialog msg for deleting
	UserState            *UserStatesMaps
	UserStateMenuMsg     *StoredMessage   // inline menu for editing and deleteing
	UserStateDialogMsgs  []*StoredMessage // dialog msg for deleting
	CallbackData         string
	CallbackText         string
	MissionAddMenuMsg    *StoredMessage   // inline menu for editing and deleteing
	MissionAddDialogMsgs []*StoredMessage // dialog msg for deleting
}

type DialogMode int

const (
	ModeMission DialogMode = 1 + iota
	ModeMissionAdd
	ModeMissionAddCallback
	ModeMissionEdit
	ModeMissionList
	ModeMissionDelCallback
	// ModeRequestInfo
	ModeStart
	ModeClear
	ModeInline
	ModeEditUserStates
	ModeEditUserStatesCallback
)

type ModeStage int

const (
	StageEmpty ModeStage = 1 + iota
	StageMissionPlace
	StageMissionStart
	StageMissionEnd
	StageMissionEdit
	StageMissionDel
	StageMissionList
	StageMissionSave
	StageMissionExit
	StageUserStateIntro
	StageUserStateExam
)

func (k ModeStage) String() string {
	switch k {
	case StageEmpty:
		return "StageEmpty"
	case StageMissionPlace:
		return "StageMissionPlace"
	case StageMissionStart:
		return "StageMissionStart"
	case StageMissionEnd:
		return "StageMissionEnd"
	case StageUserStateIntro:
		return "StageUserStateIntro"
	case StageUserStateExam:
		return "StageUserStateExam"
	case StageMissionEdit:
		return "StageMissionEdit"
	case StageMissionDel:
		return "StageMissionDel"
	case StageMissionList:
		return "StageMissionList"
	case StageMissionSave:
		return "StageMissionSave"
	case StageMissionExit:
		return "StageMissionExit"
	default:
		return "none"
	}
}

func (k DialogMode) String() string {
	switch k {
	case ModeMission:
		return "ModeMission"
	case ModeMissionAdd:
		return "ModeMissionAdd"
	case ModeMissionAddCallback:
		return "ModeMissionAddCallback"
	case ModeMissionEdit:
		return "ModeMissionEdit"
	case ModeMissionList:
		return "ModeMissionList"
	case ModeMissionDelCallback:
		return "ModeMissionDelCallback"
	// case ModeRequestInfo:
	// 	return "ModeRequestInfo"
	case ModeStart:
		return "ModeStart"
	case ModeClear:
		return "ModeClear"
	case ModeInline:
		return "ModeInline"
	case ModeEditUserStates:
		return "ModeEditUserStates"
	case ModeEditUserStatesCallback:
		return "ModeEditUserStatesCallback"
	default:
		return "none"
	}
}
