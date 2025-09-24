package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName = "access-token"
	// Just for the demo purpose, I declared a secret here. In the real-world application, you might need to get it from the env variables.
	jwtSecretKey           = "some-secret-key"
	refreshTokenCookieName = "refresh-token"
	jwtRefreshSecretKey    = "some-refresh-secret-key"
	domainCookies          = "/"
)

type Auth struct {
	app entity.Application
}

func New(app entity.Application) *Auth {
	return &Auth{
		app: app,
	}
}

func (a *Auth) GetJWTSecret() string {
	return jwtSecretKey
}

func (a *Auth) GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time.
type Claims struct {
	Login string `json:"name"`
	jwt.StandardClaims
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
// func GenerateTokensAndSetCookies(user *user.User, c echo.Context) error {
// 	accessToken, exp, err := generateAccessToken(user)
// 	if err != nil {
// 		return err
// 	}

// 	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
// 	setUserCookie(user, exp, c)

//		return nil
//	}
func (a *Auth) GenerateTokensAndSetCookies(user *entity.User, c echo.Context) error {
	defer a.RecoverReset("GenerateTokensAndSetCookies", c)
	accessToken, exp, err := a.generateAccessToken(user)
	if err != nil {
		return err
	}

	a.setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	a.setUserCookie(user, exp, c)
	// We generate here a new refresh token and saving it to the cookie.
	refreshToken, exp, err := a.generateRefreshToken(user)
	if err != nil {
		return err
	}
	a.setTokenCookie(refreshTokenCookieName, refreshToken, exp, c)

	return nil
}

func (a *Auth) generateRefreshToken(user *entity.User) (string, time.Time, error) {
	// Declare the expiration time of the token - 24 hours.
	// expirationTime := time.Now().Add(24 * time.Hour)
	// !!!DEBUG
	expirationTime := time.Now().Add(24 * time.Hour)
	// expirationTime := time.Now().Add(120 * time.Minute)

	return a.generateToken(user, expirationTime, []byte(a.GetRefreshJWTSecret()))
}

func (a *Auth) generateAccessToken(user *entity.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	// expirationTime := time.Now().Add(24 * time.Hour)
	expirationTime := time.Now().Add(60 * time.Minute)

	return a.generateToken(user, expirationTime, []byte(a.GetJWTSecret()))
}

// Pay attention to this function. It holds the main JWT token generation logic.
func (a *Auth) generateToken(user *entity.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &Claims{
		Login: user.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// Here we are creating a new cookie, which will store the valid JWT token.
func (a *Auth) setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Expires = expiration
	cookie.Path = domainCookies
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// Purpose of this cookie is to store the user's name.
func (a *Auth) setUserCookie(user *entity.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiration
	cookie.Path = domainCookies
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user try to access a protected path.
func (a *Auth) JWTErrorChecker(c echo.Context, err error) error {
	// Redirects to the signIn form.
	// return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
	return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
}

func (a *Auth) TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer a.RecoverReset("TokenRefresherMiddleware", c)
		// If the user is not authenticated (no user token data in the context), don't do anything.
		if c.Get("user") == nil {
			return next(c)
		}
		// Gets user token from the context.
		if u, ok := c.Get("user").(*jwt.Token); ok {
			claims := u.Claims.(*Claims)
			// We ensure that a new token is not issued until enough time has elapsed.
			// In this case, a new token will only be issued if the old token is within
			// 15 mins of expiry.
			if tu := time.Unix(claims.ExpiresAt, 0); time.Until(tu) < 15*time.Minute {
				// Gets the refresh token from the cookie.
				rc, err := c.Cookie(refreshTokenCookieName)
				if err == nil && rc != nil {
					// Parses token and checks if it valid.
					tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
						return []byte(a.GetRefreshJWTSecret()), nil
					})
					if err != nil {
						if err == jwt.ErrSignatureInvalid {
							c.Response().Writer.WriteHeader(http.StatusUnauthorized)
						}
					}
					if tkn != nil && tkn.Valid {
						// If everything is good, update tokens.
						_ = a.GenerateTokensAndSetCookies(&entity.User{
							Login: claims.Login,
						}, c)
					}
				}
			}
		}

		return next(c)
	}
}

func (a *Auth) RecoverReset(str string, c echo.Context) {
	if r := recover(); r != nil {
		fmt.Printf("RecoverLog: %v %v %v", str, time.Now(), r)
		a.ResetCookies(c)
	}
}

func (a *Auth) ResetCookies(c echo.Context) {
	emptyUser := &http.Cookie{
		Name:    "user",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-1e9),
		MaxAge:  -1,
	}
	emptyAccessToken := &http.Cookie{
		Name:     "access-token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1e9),
		MaxAge:   -1,
		HttpOnly: true,
	}
	emptyRefreshToken := &http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1e9),
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCookie(emptyUser)
	c.SetCookie(emptyAccessToken)
	c.SetCookie(emptyRefreshToken)
}
