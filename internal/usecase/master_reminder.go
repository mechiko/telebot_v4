package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
	"github.com/samber/lo"
)

func (uc *usecase) MasterReminder() error {
	defer uc.App.GetRecovery().RecoverLog("usecase:MasterReminder")
	// if err := uc.ApiClearAll(); err != nil {
	// 	uc.App.GetLogger().Errorf("usecase:MasterReminder %s", err.Error())
	// }
	if masters, err := uc.App.GetRepo().GetMasters().GetAll(); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		if len(masters.Items) > 0 {
			// обрабатываем мастеров
			if adminInfo, err := uc.getMasterReminderAdmin(); err != nil {
				uc.App.GetLogger().Errorf("usecase:MasterReminder %s", err.Error())
			} else {
				if !uc.isEmptyInfo(adminInfo) {
					for _, master := range masters.Items {
						if master.IsAdmin {
							adminInfo.UserIdent = master.UserIdent
							adminInfo.UserId = master.UserId
							if strTemplate, err := template.NewTemplate(uc.App, "").ReminderMasterInfo(adminInfo); err != nil {
								uc.App.GetLogger().Errorf("usecase:MasterReminder %w", err)
								continue
							} else {
								if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
									uc.App.GetLogger().Errorf("usecase:MasterReminder %w", err)
								}
								if err := uc.App.GetBot().SendMessageUser(master.UserId, strTemplate, "HTML"); err != nil {
									uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
								}
							}
						} else {
							ident := master.UserIdent
							resultNoAdmin := &entity.MasterReminderInfo{
								UserIdent: ident,
								UserId:    master.UserId,
								MissionDay0: lo.Filter(adminInfo.MissionDay0, func(x *entity.MissionUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								MissionDay3: lo.Filter(adminInfo.MissionDay3, func(x *entity.MissionUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								MissionDay30: lo.Filter(adminInfo.MissionDay30, func(x *entity.MissionUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								MissionBad: lo.Filter(adminInfo.MissionBad, func(x *entity.MissionUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								ExamenDay0: lo.Filter(adminInfo.ExamenDay0, func(x *entity.ExamenUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								ExamenDay3: lo.Filter(adminInfo.ExamenDay3, func(x *entity.ExamenUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								ExamenDay30: lo.Filter(adminInfo.ExamenDay30, func(x *entity.ExamenUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								ExamenBad: lo.Filter(adminInfo.ExamenBad, func(x *entity.ExamenUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
								ExamenEmpty: lo.Filter(adminInfo.ExamenEmpty, func(x *entity.ExamenUser, index int) bool {
									return strings.Contains(x.Masters, ident) || x.UserIdent == ident
								}),
							}
							if !uc.isEmptyInfo(resultNoAdmin) {
								if strTemplate, err := template.NewTemplate(uc.App, "").ReminderMasterInfo(resultNoAdmin); err != nil {
									uc.App.GetLogger().Errorf("usecase:MasterReminder %w", err)
									continue
								} else {
									if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
										uc.App.GetLogger().Errorf("usecase:MasterReminder %w", err)
									}
									if err := uc.App.GetBot().SendMessageUser(master.UserId, strTemplate, "HTML"); err != nil {
										uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (uc *usecase) isEmptyInfo(info *entity.MasterReminderInfo) bool {
	length := len(info.MissionDay0)
	length += len(info.MissionDay3)
	length += len(info.MissionDay30)
	length += len(info.ExamenDay0)
	length += len(info.ExamenDay3)
	length += len(info.ExamenDay30)
	length += len(info.ExamenBad)
	length += len(info.ExamenEmpty)
	return length == 0
}

func (uc *usecase) getMasterReminderAdmin() (*entity.MasterReminderInfo, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:getMasterReminderAdmin")
	result := &entity.MasterReminderInfo{
		MissionDay0:   make([]*entity.MissionUser, 0),
		MissionDay3:   make([]*entity.MissionUser, 0),
		MissionDay30:  make([]*entity.MissionUser, 0),
		MissionBad:    make([]*entity.MissionUser, 0),
		ExamenDay0:    make([]*entity.ExamenUser, 0),
		ExamenDay3:    make([]*entity.ExamenUser, 0),
		ExamenDay30:   make([]*entity.ExamenUser, 0),
		ExamenDayOt60: make([]*entity.ExamenUser, 0),
		ExamenEmpty:   make([]*entity.ExamenUser, 0),
		ExamenBad:     make([]*entity.ExamenUser, 0),
	}
	l := uc.App.GetConfiguration().Layouts.TimeLayoutDay
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(loc)
	nowDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if missions, err := uc.App.GetRepo().GetViewMissions().GetAllActive(); err != nil {
		uc.App.GetLogger().Errorf("usecase:getMasterReminderAdmin %s", err)
	} else {
		for _, mm := range missions.Items {
			if mm.Start != "" {
				if startMission, err := time.ParseInLocation(l, mm.Start, loc); err != nil {
					uc.App.GetLogger().Errorf("usecase:getMasterReminderAdmin %s", err)
					result.MissionBad = append(result.MissionBad, mm)
				} else {
					days := int(startMission.Sub(nowDate).Hours() / 24)
					switch days {
					case 0:
						result.MissionDay0 = append(result.MissionDay0, mm)
					case 3:
						result.MissionDay3 = append(result.MissionDay3, mm)
					case 30:
						result.MissionDay30 = append(result.MissionDay30, mm)
					}
				}
			} else {
				result.MissionBad = append(result.MissionBad, mm)
			}
		}
	}
	// берем экзамены через вид преврещаем ее в массив структур и сортируем
	if examens, err := uc.App.GetRepo().GetViewExamens().GetAll(); err != nil {
		uc.App.GetLogger().Errorf("usecase:getMasterReminderAdmin %s", err)
	} else {
		for _, ee := range examens.Items {
			if ee.Date != "" {
				if startExamen, err := time.ParseInLocation(l, ee.Date, loc); err != nil {
					uc.App.GetLogger().Errorf("usecase:getMasterReminderAdmin %s", err)
					result.ExamenBad = append(result.ExamenBad, ee)
				} else {
					days := int(startExamen.Sub(nowDate).Hours() / 24)
					switch days {
					case 0:
						result.ExamenDay0 = append(result.ExamenDay0, ee)
					case 3:
						result.ExamenDay3 = append(result.ExamenDay3, ee)
					case 30:
						result.ExamenDay30 = append(result.ExamenDay30, ee)
					case 60:
						if ee.ExamenKey == "ОТ" {
							result.ExamenDayOt60 = append(result.ExamenDayOt60, ee)
						}
					}
				}
			} else {
				result.ExamenEmpty = append(result.ExamenEmpty, ee)
			}
		}
	}
	return result, nil
}
