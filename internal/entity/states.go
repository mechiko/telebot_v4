package entity

import (
	"context"
	"sync"
)

// интерфейс работы с хранилищем состояний
type IStates interface {
	GetMap() MapStates
	GetState(int64) *StateDialog
	Clear()
	Run(context.Context) error
	DeleteState(int64)
}

type MapStates map[int64]*StateDialog

type States struct {
	app   Application
	In    chan *StateDialog
	mutex sync.Mutex
	buf   MapStates
}

// var _ IHistory = &History{}
var _ IStates = (*States)(nil)

// New constructs and returns a new Queue.
func NewStates(app Application) *States {
	return &States{
		app: app,
		In:  make(chan *StateDialog, 1),
		buf: make(MapStates),
	}
}

// запускаем обработку канала по закрытию канала цикл завершится
// https://stackoverflow.com/questions/45786042/go-channel-infinite-loop
func (h *States) Run(ctx context.Context) error {
	go func() {
		for s := range h.In {
			// "History chan_len: %v len_buf: %v push:%s\n", len(h.In), len(h.buf), s
			h.app.GetLogger().Debugf("StateDialogs push %v ", s.ChatId)
			h.push(s)
		}
	}()
	<-ctx.Done()
	h.app.GetLogger().Debugf("StateDialogs receive context done")
	close(h.In)
	return nil
}

func (h *States) push(s *StateDialog) {
	h.buf[s.ChatId] = s
}

func (h *States) GetMap() MapStates {
	return h.buf
}

func (h *States) Clear() {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.buf = make(MapStates)
}

func (h *States) GetState(u int64) *StateDialog {
	if state, ok := h.buf[u]; ok {
		return state
	}
	return nil
}

func (h *States) DeleteState(u int64) {
	delete(h.buf, u)
}
