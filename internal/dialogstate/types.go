package dialogstate

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

// создаем прокси на структуру состояния диалога, чтобы организовать в новом модуле все действия
// и избежать перекрестных ссылок когда импортировать будем тип StateDialog то там то там
type dialogState struct {
	*entity.StateDialog
	App    entity.Application
	States *entity.States
}

func New(app entity.Application, states *entity.States, s *entity.StateDialog) entity.DialogState {
	return &dialogState{
		StateDialog: s,
		App:         app,
		States:      states,
	}
}
