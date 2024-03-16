package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"shop-backend/api/response"
	"shop-backend/model/user"
	"shop-backend/security/accesstoken"
	"time"
)

type loginBody struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func LoginPost(e echo.Context) error {
	body := &loginBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "invalid information")
	}

	if body.Email == "" && body.Phone == "" {
		return response.BadRequestError(e, "information missing")
	}

	u, err := user.GetUserByEmailORPhone(body.Email, body.Phone)
	if err != nil {
		if errors.Is(err, user.ErrorUserNotFound) {
			return response.OtherErrors(e, response.ErrorWrongCredentials, "")
		}
		return response.ServerError(e, err, "")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(body.Password))
	if err != nil {
		return response.OtherErrors(e, response.ErrorWrongCredentials, "")
	}

	token, err := accesstoken.GenerateToken(u.ID)
	if err != nil {
		return response.ServerError(e, err, "server error")
	}

	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = token
	cookie.Expires = time.Now().Add(48 * time.Hour)
	e.SetCookie(cookie)
	return response.Success(e)

}
