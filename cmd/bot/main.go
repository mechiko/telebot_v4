package main

import (
	"time"

	"github.com/mechiko/telebot_v4/internal/cmdapp"
	"github.com/mechiko/telebot_v4/pkg/zaplog"
)

func main() {
	// Run
	exit, err := cmdapp.Run()
	zaplog.Logger.Sugar().Debugf("main() for Run exit code [%v] error [%v]", exit, err)
	time.Sleep(2 * time.Second)
	zaplog.Logger.Sugar().Debugf("END OF GAME!\n")
}
