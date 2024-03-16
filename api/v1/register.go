package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"shop-backend/api/response"
	"shop-backend/model/user"
	"shop-backend/security/accesstoken"
	"time"
)

type register struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func RegisterPOST(e echo.Context) error {
	body := &register{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "invalid information")
	}

	if body.Password == "" {
		return response.BadRequestError(e, "invalid password")
	}

	if body.Email == "" {
		return response.BadRequestError(e, "invalid email")
	}

	if body.Phone == "" || len(body.Phone) != 10 {
		return response.BadRequestError(e, "invalid phone number")
	}

	_, err := mail.ParseAddress(body.Email)
	if err != nil {
		return response.BadRequestError(e, "wrong email provided")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return response.ServerError(e, err, "server error")
	}

	//add for now email verification later
	u := user.New()
	u.ID = primitive.NewObjectID().Hex()
	u.Password = string(pass)
	u.Phone = body.Phone
	u.IsVerified = false
	u.Email = body.Email
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	err = u.Create()
	if err != nil {
		if errors.Is(err, user.ErrorEmailTaken) {
			return response.OtherErrors(e, response.ErrorEmailTaken, "")
		}
		if errors.Is(err, user.ErrorPhoneNumberTaken) {
			return response.OtherErrors(e, response.ErrorPhoneTaken, "")
		}
		return response.ServerError(e, err, "server error")
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
