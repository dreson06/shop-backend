package user

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop-backend/data"
	"strings"
	"time"
)

var ErrorUsernameTaken = errors.New("username taken")
var ErrorPhoneNumberTaken = errors.New("phone number taken")
var ErrorEmailTaken = errors.New("email taken")

type User struct {
	ID         string    `db:"id" json:"id"`
	Username   string    `db:"username" json:"username"`
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	Phone      string    `db:"phone" json:"phone"`
	IsVerified bool      `db:"is_verified" json:"is_verified"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

func New() *User {
	return &User{}
}

func (u *User) Create() error {
	if u.ID == "" {
		primitive.NewObjectID().Hex()
	}

	if u.Username == "" || u.Password == "" {
		return errors.New("some information missing")
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = u.CreatedAt

	db := data.DB()
	_, err := db.NamedExec("INSERT INTO _user (id, username, email, password, phone,is_verified, created_at, updated_at) VALUES (:id,:username,:email,:password,:phone,:is_verified,:created_at,:updated_at)", u)

	if err != nil {
		if strings.Contains(err.Error(), "phone_key") {
			return ErrorPhoneNumberTaken
		}
		if strings.Contains(err.Error(), "email_key") {
			return ErrorEmailTaken
		}
		if strings.Contains(err.Error(), "username_key") {
			return ErrorUsernameTaken
		}
		return err
	}
	return nil
}
