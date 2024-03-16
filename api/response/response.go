package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-backend/utils/logger"
)

type Status int

const (
	StatusOK           Status = 200
	StatusBadRequest   Status = 400
	StatusUnauthorized Status = 401
	StatusServerError  Status = 500

	ErrorWrongCredentials Status = 4001
	ErrorEmailTaken       Status = 4002
	ErrorPhoneTaken       Status = 4003
)

type response struct {
	Status Status `json:"status"`
	Data   any    `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

func JSON(e echo.Context, data any) error {
	r := &response{Status: StatusOK, Data: data}
	return e.JSON(http.StatusOK, r)
}

func OtherErrors(e echo.Context, code Status, msg string) error {
	r := &response{Status: code, Error: msg}
	logger.L.With("item", "api", "code", code).Errorln(msg)
	return e.JSON(http.StatusOK, r)
}

func UnauthorizedError(e echo.Context) error {
	return OtherErrors(e, StatusUnauthorized, "")
}

func BadRequestError(e echo.Context, msg string) error {
	r := &response{Status: StatusBadRequest, Error: msg}
	logger.L.Errorln(msg)
	return e.JSON(http.StatusOK, r)
}

func ServerError(e echo.Context, err error, msg string) error {
	r := &response{Status: StatusServerError, Error: msg}
	logger.L.Errorln(err)
	return e.JSON(http.StatusOK, r)
}

func Success(e echo.Context) error {
	r := &response{Status: StatusOK, Data: map[string]interface{}{"success": true}}
	return e.JSON(http.StatusOK, r)
}
