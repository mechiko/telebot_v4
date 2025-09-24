package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *AdminRoutes) GetMasterUsers(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("GetMasterUsers")
	if list, err := r.app.GetRepo().GetMasters().GetAll(); err != nil {
		r.app.GetLogger().Errorf("admin:GetMasterUsers %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, list); err != nil {
			r.app.GetLogger().Errorf("admin:GetMasterUsers %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}
