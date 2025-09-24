// Package v1 implements routing paths. Each services in own file.
package v1

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mechiko/telebot_v4/internal/controller/http/v1/admin"
	"github.com/mechiko/telebot_v4/internal/controller/http/v1/auth"
	v1db "github.com/mechiko/telebot_v4/internal/controller/http/v1/db"
	"github.com/mechiko/telebot_v4/internal/controller/http/v1/user"
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
	"github.com/mechiko/telebot_v4/pkg/zaplog"
	"golang.org/x/crypto/bcrypt"
)

type V1Routes struct {
	app  entity.Application
	su   entity.UseCase
	auth *auth.Auth
}

//go:embed login.html
var LoginHTML string

func New(handler *echo.Echo, app entity.Application) *V1Routes {
	// Options
	// handler.Use(gin.Logger())
	// handler.Use(gin.Recovery())
	r := &V1Routes{
		app:  app,
		su:   usecase.New(app),
		auth: auth.New(app),
	}

	// K8s probe
	handler.GET("/healthz", func(c echo.Context) error {
		fmt.Println("GET(/healthz, func(c echo.Context)")
		return c.String(200, "ok")
	})
	handler.Pre(middleware.RemoveTrailingSlash())
	// Routers
	// http://localhost:3600/v1/db/name
	h := handler.Group("/v1")
	{
		h.GET("/user/signin", r.SignInForm).Name = "userSignInForm"
		h.POST("/user/signin", r.SignIn)
		h.GET("/user/signout", r.SignOut)
		h.GET("", r.Home)
	}

	v1db.New(h, app)
	admin.New(h, app)
	return r
}

func (r *V1Routes) Home(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("Home")
	// for _, cookie := range c.Cookies() {
	// 	fmt.Println(cookie.Name)
	// 	fmt.Println(cookie.Value)
	// }
	// fp := path.Join("templates", "signIn.html")
	return c.String(http.StatusOK, "HOME")
}

// SignInForm responsible for signIn Form rendering.
func (r *V1Routes) SignInForm(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("SignInForm")
	r.auth.ResetCookies(c)
	tmpl, err := template.New("login").Parse(LoginHTML)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := tmpl.Execute(c.Response().Writer, nil); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// SignIn will be executed after SignInForm submission.
func (r *V1Routes) SignIn(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("SignIn")
	// Load our "test" user.
	// Initiate a new User struct.
	r.auth.ResetCookies(c)
	u := new(user.User)
	// Parse the submitted data and fill the User struct with the data from the SignIn form.
	if err := c.Bind(u); err != nil {
		r.app.GetLogger().Error("auth:SignIn", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	r.app.GetLogger().Infof("Attempt Login: %s Pass: %s", u.Username, u.Password)
	requestIP := c.RealIP()
	zaplog.AuthSugar.Infof("Attempt Login: %s Pass: %s from: %s", u.Username, u.Password, requestIP)

	storedUser, err := r.app.GetRepo().GetUsers().GetByLoginAdmin(u.Username)
	if err != nil {
		r.app.GetLogger().Errorf("Attempt SignIn check Login: %s error: %s", u.Username, err.Error())
		zaplog.AuthSugar.Infof("Attempt SignIn check Login: %s error: %s", u.Username, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("user login or password wrong"))
	}
	// Compare the stored hashed password, with the hashed version of the password that was received.
	dp, err := hex.DecodeString(storedUser.Passwd)
	if err != nil {
		r.app.GetLogger().Errorf("Attempt SignIn passwd decode Login: %s error: %s", u.Username, err.Error())
		zaplog.AuthSugar.Infof("Attempt SignIn passwd decode Login: %s error: %s", u.Username, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("user login or password wrong"))
	}
	if err := bcrypt.CompareHashAndPassword(dp, []byte(u.Password)); err != nil {
		// If the two passwords don't match, return a 401 status.
		r.app.GetLogger().Errorf("Attempt SignIn passwd compare Login: %s error: %s", u.Username, err.Error())
		zaplog.AuthSugar.Infof("Attempt SignIn passwd compare Login: %s error: %s", u.Username, err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, "user login or password wrong")
	}
	// If password is correct, generate tokens and set cookies.
	if err := r.auth.GenerateTokensAndSetCookies(storedUser, c); err != nil {
		r.app.GetLogger().Errorf("Attempt SignIn generate token Login: %s error: %s", u.Username, err.Error())
		zaplog.AuthSugar.Infof("Attempt SignIn generate token Login: %s error: %s", u.Username, err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	zaplog.AuthSugar.Infof("Success Login: %s Pass: %s from: %s", u.Username, u.Password, requestIP)
	return c.String(http.StatusOK, "ok")
	// return c.Redirect(http.StatusMovedPermanently, "/v1/admin")
}

func (r *V1Routes) SignOut(c echo.Context) error {
	defer r.app.GetRecovery().RecoverLog("SignOut")
	for _, cookie := range c.Cookies() {
		switch cookie.Name {
		case "user":
			zaplog.AuthSugar.Infof("Attempt SignOut Login: %s", cookie.Value)
		}
	}
	r.auth.ResetCookies(c)
	return c.String(http.StatusOK, "logout")
}
