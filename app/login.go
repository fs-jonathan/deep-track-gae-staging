package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type LineUser struct {
	UserId string `json:"lineUserId"`
}

type FirebaseUser struct {
	UserId string `json:"userid"`
}

func init() {
  e.POST("/loginLiff", liffLogin)
  e.POST("/loginReact", reactLogin)
}

func liffLogin(c echo.Context) error {
	// Received Line UserId
	user := new(LineUser)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println(user.UserId)

	message := Message{0}
	return c.JSON(http.StatusOK, message)
}

func reactLogin(c echo.Context) error {
	// Received Firebase UserId
	user := new(FirebaseUser)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println(user.UserId)

	return c.NoContent(http.StatusOK)
}
