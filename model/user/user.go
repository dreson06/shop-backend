package user

import (
	"database/sql"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop-backend/data"
	"strings"
	"time"
)

var ErrorPhoneNumberTaken = errors.New("phone number taken")
var ErrorEmailTaken = errors.New("email taken")
var ErrorUserNotFound = errors.New("user not found")

type User struct {
	ID         string    `db:"id" json:"id"`
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

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = u.CreatedAt

	db := data.DB()
	_, err := db.NamedExec("INSERT INTO _user (id, email, password, phone,is_verified, created_at, updated_at) VALUES (:id,:email,:password,:phone,:is_verified,:created_at,:updated_at)", u)

	if err != nil {
		if strings.Contains(err.Error(), "phone_key") {
			return ErrorPhoneNumberTaken
		}
		if strings.Contains(err.Error(), "email_key") {
			return ErrorEmailTaken
		}
		return err
	}
	return nil
}

func GetUserByEmailORPhone(email, phone string) (*User, error) {
	db := data.DB()
	u := New()

	err := db.Get(u, "SELECT * FROM _user WHERE email = $1 OR phone = $2", email, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}
	return u, nil
}
