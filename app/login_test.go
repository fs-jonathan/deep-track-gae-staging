package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	lineUserJSON = `{"message":0}`
)

func TestLiffLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/loginLiff", strings.NewReader(lineUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, liffLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, lineUserJSON, rec.Body.String())
	}
}

func TestReactLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/loginReact", strings.NewReader(lineUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, reactLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}
