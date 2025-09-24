package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *AdminRoutes) GetAllExamen(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("admin:GetAllExamen")
	if all, err := r.app.GetRepo().GetViewExamens().GetAll(); err != nil {
		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, all); err != nil {
			r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	// if all, err := usecase.New(r.app).ApiGetAllExamens(""); err != nil {
	// 	r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// } else {
	// 	if err := c.JSON(http.StatusOK, all); err != nil {
	// 		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
	// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}
	// }
	return nil
}

func (r *AdminRoutes) GetAllExamenMaster(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("admin:GetAllExamen")
	// master := c.Param("master")

	if all, err := r.app.GetRepo().GetViewExamens().GetAll(); err != nil {
		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if err := c.JSON(http.StatusOK, all); err != nil {
			r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	// if all, err := usecase.New(r.app).ApiGetAllExamens(master); err != nil {
	// 	r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// } else {
	// 	if err := c.JSON(http.StatusOK, all); err != nil {
	// 		r.app.GetLogger().Errorf("admin:GetAllExamen %s", err.Error())
	// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}
	// }
	return nil
}
