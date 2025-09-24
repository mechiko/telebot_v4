package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"

	"github.com/mechiko/telebot_v4/internal/bot/middle"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/layout"
)

type Bot struct {
	*telebot.Bot
	L      *layout.Layout
	App    entity.Application
	States *entity.States // хранилище состояний диалогов с пользователями
	// UserStates entity.UsersStatesMap // все состояния пользователей из БД при инициализации бота и по мере работы обновляется
	// Inline bool
	// MsgUserState map[int64]*tele.Message
}

func (b *Bot) GetStates() *entity.States {
	return b.States
}

func New(app entity.Application, path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}
	states := entity.NewStates(app)
	sets := lt.Settings()
	sets.OnError = func(err error, c telebot.Context) {
		app.GetLogger().Errorf("telebot %v %v", c.Sender().Recipient(), err)
	}
	sets.Token = app.GetConfiguration().Token
	sets.Poller = &telebot.LongPoller{Timeout: 10 * time.Second}
	sets.Verbose = false
	// .AllowedUpdates = []string{"callback_query", "message"}

	b, err := telebot.NewBot(sets)
	if err != nil {
		return nil, err
	}
	params := telebot.CommandScope{Type: telebot.CommandScopeAllPrivateChats}
	cmds := lt.Commands()
	if cmds != nil {
		if err := b.SetCommands(cmds, params, "ru"); err != nil {
			return nil, err
		}
	}
	// для группового чата это и второе
	groupScope := &telebot.CommandScope{Type: telebot.CommandScopeChatMember}
	// adminScope := &telebot.CommandScope{Type: telebot.CommandScopeChatAdmin}
	cmdsGroup := make([]telebot.Command, 0)
	commandStart := telebot.Command{Text: "start", Description: "привет"}
	cmdsGroup = append(cmdsGroup, commandStart)
	if err := b.SetCommands(cmdsGroup, groupScope, "ru"); err != nil {
		return nil, err
	}

	// if getCommands, err := b.Commands(groupScope, "ru"); err != nil {
	// 	return nil, err
	// } else {
	// 	app.GetLogger().Debugf("group ru %v", getCommands)
	// }
	// if getCommands, err := b.Commands(groupScope); err != nil {
	// 	return nil, err
	// } else {
	// 	app.GetLogger().Debugf("group default %v", getCommands)
	// }
	// if getCommands, err := b.Commands(params, "ru"); err != nil {
	// 	return nil, err
	// } else {
	// 	app.GetLogger().Debugf("private ru %v", getCommands)
	// }
	// if getCommands, err := b.Commands(params); err != nil {
	// 	return nil, err
	// } else {
	// 	app.GetLogger().Debugf("private default %v", getCommands)
	// }

	return &Bot{
		Bot:    b,
		L:      lt,
		App:    app,
		States: states,
	}, nil
}

// запуск самого бота не команды старт
func (b *Bot) Start(ctx context.Context) error {
	ctxUrls, cancel := context.WithCancel(context.Background())

	// запускаем хранилище urls в фоне иначе не будет обработка по каналу
	go func() {
		b.States.Run(ctxUrls)
	}()

	// Middlewares
	b.Use(middle.Logger(b.App))
	// обработак команд в группе для прописки пользователя в базу для работы с ним
	b.Use(middle.GroupChecker(b.App))
	// фильтр по писателю и приватному чату
	b.Use(middle.SenderChecker(b.App))
	// обработка состояний фильтрации входящих и прочее
	b.Use(middle.Handler(b.App))
	b.Use(middle.RestricterIn(b.App))
	// это для запуска Layout location
	b.Use(b.L.Middleware("en", func(r telebot.Recipient) string {
		return "ru"
	}))
	// ruCommands := b.L.CommandsLocale("ru")
	// enCommands := b.L.CommandsLocale("en")
	// b.SetCommands(ruCommands, &tele.CommandScope{Type: tele.CommandScopeAllChatAdmin}, "ru")
	// b.SetCommands(enCommands, &tele.CommandScope{Type: tele.CommandScopeAllChatAdmin}, "en")

	// Handlers
	// b.Handle("/start", b.onStart)
	// b.Handle("/hello", b.onHello)
	// b.Handle("/clear", b.onClear)
	// b.Handle(b.L.ButtonLocale("ru", "mission"), b.onMission)
	// b.Handle(b.L.ButtonLocale("ru", "missionAdd"), b.onMissionAdd)
	// b.Handle(b.L.ButtonLocale("ru", "missionList"), b.onMissionList)
	// b.Handle(b.L.ButtonLocale("ru", "missionEdit"), b.onMissionEdit)
	// b.Handle(b.L.ButtonLocale("ru", "missionDel"), b.onMissionDel)
	// b.Handle(b.L.ButtonLocale("ru", "exit"), b.onExit)
	// b.Handle(b.L.ButtonLocale("ru", "return"), b.onReturn)
	// b.Handle(b.L.ButtonLocale("ru", "missionAddPlace"), b.onMissionAddStagePlace)
	// b.Handle(b.L.ButtonLocale("ru", "missionAddStart"), b.onMissionAddStageStart)
	// b.Handle(b.L.ButtonLocale("ru", "missionAddEnd"), b.onMissionAddStageEnd)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace1"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace2"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace3"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace4"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace5"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionPlace6"), b.onMissionAddStagePlaceAny)
	// b.Handle(b.L.ButtonLocale("ru", "missionAddSave"), b.onMissionAddSave)
	// b.Handle(b.L.ButtonLocale("ru", "requestInfo"), b.onRequestInfo)
	// b.Handle(b.L.Callback("intro"), b.onCallbackIntro)
	// b.Handle(b.L.Callback("exam1"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam2"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam3"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam4"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam5"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam6"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exam7"), b.onCallbackExam)
	// b.Handle(b.L.Callback("exitState"), b.onCallbackExit)
	// b.Handle(tele.OnContact, b.onContact)
	// b.Handle(tele.OnLocation, b.onContact)

	b.Handle(telebot.OnText, b.Router)
	// если есть калбэк на кнопку то общий обработчик сработает на необработанные колбэки
	b.Handle(telebot.OnCallback, b.routerCallback)
	b.App.GetLogger().Infof("run:bot started")
	go func() {
		<-ctx.Done()
		b.App.GetLogger().Debugf("bot:start receive group context done")
		cancel()
		b.Stop()
	}()
	b.Bot.Start()
	return fmt.Errorf("stop bot")
}
