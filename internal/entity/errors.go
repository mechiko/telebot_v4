package entity

import (
	"errors"
	"fmt"
	"time"
)

// для таких ошибок работает метод errors.Is(err, ErrRuleIsTimeOut)
var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given param is not valid")
	// мои придуманные ошибки
	ErrNotPingUtm             = errors.New("utm host:port not ping")
	ErrRuleIsTimeOut          = errors.New("правило ждет таймаут")
	ErrClearBackgroundTasks   = errors.New("остановка фоновых задач")
	ErrBackgroundTasksPresent = errors.New("задача для этого запроса уже добавлена")
	ErrXMLResponceZeroValue   = errors.New("ответ пустой")
	ErrNoItemsForTask         = errors.New("все данные есть, отстутствует необходимость обновления")
	ErrAppRestart             = errors.New("перезапуск приложения")
	ErrAppShutdown            = errors.New("прерывание приложения")
	ErrAppTypeDelete          = errors.New("ошибка удаления приложения есть связи")
	ErrAppTypeBadEmpty        = errors.New("ошибка данных приложения они пустые")
	ErrClientBadEmpty         = errors.New("ошибка данных клиента они пустые")
	ErrLicenseTypeDelete      = errors.New("ошибка удаления типа лицензии есть связи")
	ErrLicenseTypeEmpty       = errors.New("ошибка типа лицензии пустой")
	ErrLicenseTypeReserved    = errors.New("ошибка тип лицензии 'permanent' зарезервирован")
	ErrIdentityTypeDelete     = errors.New("ошибка удаления типа идентификатора есть связи")
	ErrIdentityTypeEmpty      = errors.New("ошибка типа идентификатора пустой")
	ErrIdentityTypeReserved   = errors.New("ошибка тип идентификатора зарезервирован")
	ErrClientDelete           = errors.New("ошибка удаления клиента есть связи")
	ErrClientEmpty            = errors.New("ошибка клиент пустой")
	ErrRegistryIdentityDelete = errors.New("ошибка удаления идентификатора клиента есть связи")
	ErrRegistryIdentityEmpty  = errors.New("ошибка идентификатор пустой")
	ErrActiveLicenseNotFound  = errors.New("ошибка нет активной лицензии")
)

//	для таких ошибок надо if perr, ok := err.(*RuleIsTimeOutError); ok {
//		TimeOut  int     для правила установлен таймаут в секундах
//		TimeLeft int     для правила осталось секунд ожидания
//		Msg      string  сообщение ошибки
type RuleIsTimeOutError struct {
	TimeOut  int
	TimeWait time.Duration
}

func (e *RuleIsTimeOutError) Error() string {
	return fmt.Sprintf("правило ждет тайм аут %v осталось %s", e.TimeOut, e.TimeWait)
}

type UTMOutError struct {
	Err error
	Doc string
}

func (e *UTMOutError) Error() string {
	return fmt.Sprintf("ошибка УТМ  %v получено: '%v'", e.Err.Error(), e.Doc)
}
