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
	message = `{"message":0}`
)

func TestJsonWriter(t *testing.T) {
  e := echo.New()
  req := httptest.NewRequest(http.MethodPost, "/getStaticJson", strings.NewReader(message))

  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, liffLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, message, rec.Body.String())
	}
}

func TestDefaultRecord(t *testing.T) {
  s0 := []Record{Record{1, "本日（現時点まで）", "", 9, 8, 7}}
  s1 := append(s0, Record{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
  s2 := append(s1, Record{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
  s3 := append(s2, Record{4, "先月", "先々月との比較", 93, 83, 72})
  expected := append(s3, Record{5, "全期間", "", 20, 4, 434})
  actual := getDefaultRecord()

  for i := range expected {
    if actual[i] != expected[i] {
        t.Errorf("got: %v\nwant: %v", actual[0], expected[0])
    }
  }
}
