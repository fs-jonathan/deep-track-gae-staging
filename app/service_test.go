package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetReport(t *testing.T) {
  e := echo.New()
  req := httptest.NewRequest(http.MethodPost, "/getReport", strings.NewReader(message))

  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, liffLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetDetail(t *testing.T) {
  e := echo.New()
  req := httptest.NewRequest(http.MethodPost, "/getDetail", strings.NewReader(message))

  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, liffLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, message, rec.Body.String())
	}
}
