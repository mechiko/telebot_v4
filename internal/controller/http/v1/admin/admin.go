package admin

import (
	"fmt"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/mechiko/telebot_v4/internal/controller/http/v1/auth"
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
)

type AdminRoutes struct {
	app entity.Application
	su  entity.UseCase
}

// http://localhost:3600/v1/admin/...
func New(handler *echo.Group, app entity.Application) {
	r := &AdminRoutes{
		app: app,
		su:  usecase.New(app),
	}
	h := handler.Group("/admin")
	a := auth.New(app)
	h.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte(a.GetJWTSecret()),
		TokenLookup:  "cookie:access-token", // "<source>:<name>"
		ErrorHandler: a.JWTErrorChecker,
	}))
	h.Use(a.TokenRefresherMiddleware)

	{
		h.GET("/ping", r.AdminPing)
		h.GET("/mission/:master", r.GetAllMissionMaster)
		h.GET("/mission/all", r.GetAllMission)
		h.GET("/examen/:master", r.GetAllExamenMaster)
		h.GET("/examen/all", r.GetAllExamen)
		h.GET("/users", r.GetUsers)
		h.GET("/telebot/users", r.GetTelebotUsers)
		h.PUT("/telebot/users", r.UpdateTelebotUser)
		h.DELETE("/telebot/users", r.DeleteTelebotUser)
		h.GET("/masters", r.GetMasterUsers)
	}
}

func (r *AdminRoutes) AdminHome(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("AdminHome")
	fmt.Println("AdminHome enter")
	for _, cookie := range c.Cookies() {
		switch cookie.Name {
		case "user":
			fmt.Printf("cookie %s:%s %v\n", cookie.Name, cookie.Value, cookie.Expires)
		case "access-token":
			fmt.Printf("cookie %s %v\n", cookie.Name, cookie.Expires)
		case "refresh-token":
			fmt.Printf("cookie %s %v\n", cookie.Name, cookie.Expires)
		}
	}
	return c.String(http.StatusOK, "ok")
}

func (r *AdminRoutes) AdminPing(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("AdminHome")
	fmt.Println("AdminPing")
	for _, cookie := range c.Cookies() {
		switch cookie.Name {
		case "user":
			fmt.Printf("cookie %s:%s %+v\n", cookie.Name, cookie.Value, cookie)
		case "access-token":
			fmt.Printf("cookie %s %v\n", cookie.Name, cookie.Expires)
		case "refresh-token":
			fmt.Printf("cookie %s %v\n", cookie.Name, cookie.Expires)
		}
	}
	return c.JSON(http.StatusOK, "pong")
}
