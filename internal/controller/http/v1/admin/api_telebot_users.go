package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *AdminRoutes) GetTelebotUsers(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("GetTelebotUsers")
	if list, err := r.su.ApiGetTelebotUsers(); err != nil {
		r.app.GetLogger().Errorf("admin:GetTelebotUsers %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, list); err != nil {
			r.app.GetLogger().Errorf("admin:GetTelebotUsers %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}

func (r *AdminRoutes) UpdateTelebotUser(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("UpdateTelebotUser")
	var user = &entity.TelebotUser{}

	if err := c.Bind(&user); err != nil {
		r.app.GetLogger().Errorf("admin:UpdateTelebotUser %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// надо сделать валидацию :)
	if user.ID == 0 {
		r.app.GetLogger().Errorf("admin:UpdateTelebotUser id empty")
		return echo.NewHTTPError(http.StatusInternalServerError, "id empty")
	}
	if user.Ident == "" {
		r.app.GetLogger().Errorf("admin:UpdateTelebotUser ident empty")
		return echo.NewHTTPError(http.StatusInternalServerError, "ident empty")
	}

	if err := r.app.GetRepo().GetTelebotUsers().Update(user); err != nil {
		r.app.GetLogger().Errorf("admin:UpdateTelebotUser %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.String(http.StatusOK, "OK"); err != nil {
			r.app.GetLogger().Errorf("admin:UpdateTelebotUser %s", err.Error())
			return err
		}
	}
	return nil
}

func (r *AdminRoutes) DeleteTelebotUser(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("DeleteTelebotUser")
	var user = &entity.TelebotUser{}

	if err := c.Bind(&user); err != nil {
		r.app.GetLogger().Errorf("admin:DeleteTelebotUser %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// надо сделать валидацию :)
	if user.ID == 0 {
		r.app.GetLogger().Errorf("admin:DeleteTelebotUser id empty")
		return echo.NewHTTPError(http.StatusInternalServerError, "id empty")
	}

	if err := r.app.GetRepo().GetTelebotUsers().Delete(user); err != nil {
		r.app.GetLogger().Errorf("admin:DeleteTelebotUser %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.String(http.StatusOK, "OK"); err != nil {
			r.app.GetLogger().Errorf("admin:DeleteTelebotUser %s", err.Error())
			return err
		}
	}
	return nil
}
