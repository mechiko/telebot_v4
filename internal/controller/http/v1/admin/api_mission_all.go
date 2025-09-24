package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *AdminRoutes) GetAllMission(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("admin:GetAllMission")

	if all, err := r.app.GetRepo().GetViewMissions().GetAllActive(); err != nil {
		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, all); err != nil {
			r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// if all, err := usecase.New(r.app).ApiGetAllMissions(""); err != nil {
	// 	r.app.GetLogger().Errorf("admin:GetAllMission %s", err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// } else {
	// 	if err := c.JSON(http.StatusOK, all); err != nil {
	// 		r.app.GetLogger().Errorf("admin:GetAllMission %s", err.Error())
	// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}
	// }
	return nil
}

func (r *AdminRoutes) GetAllMissionMaster(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("admin:GetAllMission")
	if all, err := r.app.GetRepo().GetViewMissions().GetAllActive(); err != nil {
		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, all); err != nil {
			r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	// master := c.Param("master")
	// if all, err := usecase.New(r.app).ApiGetAllMissions(master); err != nil {
	// 	r.app.GetLogger().Errorf("admin:GetAllMission %s", err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// } else {
	// 	if err := c.JSON(http.StatusOK, all); err != nil {
	// 		r.app.GetLogger().Errorf("admin:GetAllMission %s", err.Error())
	// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}
	// }
	return nil
}
