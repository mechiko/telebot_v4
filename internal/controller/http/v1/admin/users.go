package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *AdminRoutes) GetUsers(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("GetUsers")
	if list, err := r.su.ApiGetUsers(); err != nil {
		r.app.GetLogger().Errorf("admin:GetUsers %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, list); err != nil {
			r.app.GetLogger().Errorf("admin:GetUsers %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}
