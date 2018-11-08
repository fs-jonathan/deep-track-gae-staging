package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Record struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Cost     int    `json:"cost"`
	Compare  int    `json:"compare"`
	Rate     int    `json:"rate"`
}

type Message struct {
	Message int `json:"message"`
}

type LineUser struct {
	UserId string `json:"lineUserId"`
}

type FirebaseUser struct {
	UserId string `json:"userid"`
}

func init() {
	// 検証用
	// e.GET("/app", jsonWriter)

	// Route => handler
	e.POST("/loginLiff", liffLogin)
	e.POST("/loginReact", reactLogin)
	e.POST("/getJson", jsonWriter)
}

func jsonWriter(c echo.Context) error {
	// Receive message
	message := new(Message)

	if reqErr := c.Bind(message); reqErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, reqErr.Error())
	}

	// log.Println(message)
	b, err := ioutil.ReadFile("json/record" + strconv.Itoa(message.Message) + ".json")

	// Return default message on any error
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusOK, getDefaultRecord())
	}

	var records []Record
	parseErr := json.Unmarshal(b, &records)

	if parseErr != nil {
		// log.Println(parseErr)
		return c.JSON(http.StatusOK, getDefaultRecord())
	}

	return c.JSON(http.StatusOK, records)
}

func getDefaultRecord() []Record {
	s0 := []Record{Record{1, "本日（現時点まで）", "", 9, 8, 7}}
	s1 := append(s0, Record{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
	s2 := append(s1, Record{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
	s3 := append(s2, Record{4, "先月", "先々月との比較", 93, 83, 72})
	records := append(s3, Record{5, "全期間", "", 20, 4, 434})

	return records
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
