// Package app configures and runs application.
package cmdapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/telebot_v4/internal/application"
	"github.com/mechiko/telebot_v4/internal/bot"
	"github.com/mechiko/telebot_v4/internal/checkdbg"
	"github.com/mechiko/telebot_v4/internal/crontab"
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
	"github.com/mechiko/telebot_v4/pkg/htmlserver"
	"github.com/mechiko/telebot_v4/pkg/httpserver"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3"
	"golang.org/x/sync/errgroup"
)

// Run creates objects via constructors.
func Run() (int, error) {
	var err error
	var app entity.Application

	// https://habr.com/ru/company/vivid_money/blog/531822/
	if app, err = application.NewApplication(); err != nil {
		fmt.Printf("error application.NewApplication() %s\n", err.Error())
		os.Exit(1)
	}
	defer app.GetRecovery().RecoverLog("cmdapp.Run()")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.GetLogger().Infof("cmdapp:run Database.DbName:%s", app.GetConfiguration().Database.DbName)

	// инициализируем REPO тут как реализацию всего доступа к частям системы через таблицы и драйвер
	repo := sqlite3.NewRepository(app)
	app.SetRepo(repo)
	if err := repo.Start(); err != nil {
		return 0, fmt.Errorf("cmdapp:run repo.Start() %w", err)
	}
	handler := echo.New()
	httpServer := httpserver.New(app, handler, httpserver.Port(app.GetConfiguration().HostPort))
	handler2 := echo.New()
	htmlServer := htmlserver.New(app, handler2, htmlserver.Port(app.GetConfiguration().HtmlPort))

	botYamlPath := path.Join(entity.ConfigPath, "bot.yaml")

	b, err := bot.New(app, botYamlPath)
	if err != nil {
		return 0, fmt.Errorf("cmdapp:run bot.New() %w", err)
	}
	app.SetBot(b)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			app.GetLogger().Infof("получен сигнал завершения контекста группы в HTTP")
			if err := httpServer.Shutdown(); err != nil {
				app.GetLogger().Errorf("Stopped http server with error:", err)
			}
		}()
		httpServer.Start()
		// по ошибке сервера возвращаем в группу код ошибки
		return <-httpServer.Notify()
	})

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			app.GetLogger().Infof("получен сигнал завершения контекста группы в HTTP")
			if err := htmlServer.Shutdown(); err != nil {
				app.GetLogger().Errorf("Stopped http server with error:", err)
			}
		}()
		htmlServer.Start()
		// по ошибке сервера возвращаем в группу код ошибки
		return <-htmlServer.Notify()
	})

	if app.GetConfiguration().Telebot {
		group.Go(func() error {
			return b.Start(groupCtx)
		})
	}

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			sigChan <- os.Kill
		}()
		return fmt.Errorf("exit os signal %v", <-sigChan)
	})

	if app.GetConfiguration().Debug {
		group.Go(func() error {
			// вызываем проверку чего либо такого в этой процедуре
			// и без контекста завершимся быстро за проход
			return checkdbg.NewChecks(app).Run()
		})
	}

	// запускаем задачи по таймеру так, если надо просто завершить вызываем
	// cron.Shutdown()
	cronClockClear := crontab.NewClock(app, 12, 0, func() error {
		// cron := crontab.NewDuration(app, 12*time.Hour, func() error {
		defer app.GetRecovery().RecoverLog("cronDuration")
		app.GetLogger().Debugf("cron duration 12 hour send message to admin %v", time.Now().Local())
		if err := usecase.New(app).ApiClearAll(); err != nil {
			msg := fmt.Sprintf("ошибка работы очистки дат %s", err.Error())
			if err := b.SendMessageAdmin(msg, ""); err != nil {
				app.GetLogger().Errorf("cron send error %s", err.Error())
			}
		}
		msg := fmt.Sprintf("бот жив %v таймер на 12:00 сброс дат", time.Now().Local())
		if err := b.SendMessageAdmin(msg, ""); err != nil {
			app.GetLogger().Errorf("cron send error %s", err.Error())
		}
		return nil
	})

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			cronClockClear.Shutdown()
		}()
		return cronClockClear.RunClock(groupCtx)
	})

	go func() {
		<-time.After(3 * time.Second)
		app.GetLogger().Debugf("cron after 3s")
		msg := fmt.Sprintf("бот запущен в %v", time.Now().Local())
		if err := b.SendMessageAdmin(msg, ""); err != nil {
			app.GetLogger().Errorf("cron send error %s", err.Error())
		}
	}()

	// запускаем задачи по таймеру так, если надо просто завершить вызываем
	// cron.Shutdown()
	cronClockReminder := crontab.NewClock(app, 7, 0, func() error {
		defer app.GetRecovery().RecoverLog("cronClock")
		app.GetLogger().Debugf("cron CLOCK time:%s", time.Now())
		if err := usecase.New(app).UsersReminder(); err != nil {
			msg := fmt.Sprintf("ошибка работы напоминалки пользователям %s", err.Error())
			if err := b.SendMessageAdmin(msg, ""); err != nil {
				app.GetLogger().Errorf("cron send error %s", err.Error())
			}
			return nil
		}
		if err := usecase.New(app).MasterReminder(); err != nil {
			msg := fmt.Sprintf("ошибка работы напоминалки мастерам %s", err.Error())
			if err := b.SendMessageAdmin(msg, ""); err != nil {
				app.GetLogger().Errorf("cron send error %s", err.Error())
			}
			return nil
		}
		msg := fmt.Sprintf("бот жив %v таймер на 07:00 напоминалка ", time.Now().Local())
		if err := b.SendMessageAdmin(msg, ""); err != nil {
			app.GetLogger().Errorf("cron send error %s", err.Error())
		}
		return nil
	})
	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			cronClockReminder.Shutdown()
		}()
		return cronClockReminder.RunClock(groupCtx)
	})

	// go func() {
	// 	<-time.After(620 * time.Minute)
	// 	app.GetLogger().Debugf("cron CLOCK shutdown")
	// 	cronClock.Shutdown()
	// }()

	exitCode := 0
	errGroup := group.Wait()
	app.GetLogger().Infof("exit group.Wait() %v", errGroup.Error())
	return exitCode, errGroup
}
