package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"smashedbits.com/shorty/pkg/model"
)

const jwtCookieName = "jwt"

type userStorage interface {
	GetByID(ctx context.Context, userId string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	Insert(ctx context.Context, user model.User) error
}

type Authenticator struct {
	store         userStorage
	jwtSigningKey []byte
}

func NewAuthenticator(store userStorage, jwtSigningKey string) Authenticator {
	return Authenticator{
		store:         store,
		jwtSigningKey: []byte(jwtSigningKey),
	}
}

func (u Authenticator) GetUser(eCtx echo.Context) (model.User, error) {
	userId, err := u.GetUserID(eCtx)
	if err != nil {
		return model.User{}, err
	}

	ctx := eCtx.Request().Context()
	return u.store.GetByID(ctx, userId)
}

func (u Authenticator) GetUserID(eCtx echo.Context) (string, error) {
	jwtStg, err := u.GetJWTCookie(eCtx)
	if err != nil {
		return "", err
	}

	token, err := u.validateJWT(jwtStg)
	if err != nil {
		return "", err
	}

	subj, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return subj, nil
}

func (u Authenticator) CompleteUserSignIn(eCtx echo.Context, retryIfLoggedOut bool) error {
	req := eCtx.Request()
	res := eCtx.Response()
	ctx := req.Context()

	gothUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		if retryIfLoggedOut {
			gothic.BeginAuthHandler(res, req)
		}
		return err
	}

	user, err := u.upsertUser(ctx, gothUser)
	if err != nil {
		return err
	}

	u.generateJWTCookie(eCtx, user.ID)

	return nil
}

func (a Authenticator) GetJWTCookie(eCtx echo.Context) (string, error) {
	cookie, err := eCtx.Cookie(jwtCookieName)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (u Authenticator) generateJWTCookie(eCtx echo.Context, userId string) error {
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	jwtStg, err := u.generateJWT(userId, expiresAt)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     jwtCookieName,
		Value:    jwtStg,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
	}
	eCtx.SetCookie(cookie)

	return nil
}

func (u Authenticator) Logout(eCtx echo.Context) error {
	req := eCtx.Request()
	res := eCtx.Response()

	gothic.Logout(res, req)

	cookie := &http.Cookie{
		Name:     jwtCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	eCtx.SetCookie(cookie)

	return nil
}

func (u Authenticator) upsertUser(ctx context.Context, gothUser goth.User) (model.User, error) {
	user, err := u.store.GetByEmail(ctx, gothUser.Email)
	if err != nil {
		user = model.User{
			Email: gothUser.Email,
		}

		if err := u.store.Insert(ctx, user); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (u Authenticator) generateJWT(userId string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Subject:   userId,
	})
	token.Method = jwt.SigningMethodHS256

	return token.SignedString(u.jwtSigningKey)
}

func (u Authenticator) validateJWT(jwtStg string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtStg, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.jwtSigningKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
