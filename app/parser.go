package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Record struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	Subtitle string  `json:"subtitle"`
	Cost     float64 `json:"cost"`
	Compare  float64 `json:"compare"`
	Rate     float64 `json:"rate"`
}

type Message struct {
	Message int `json:"message"`
}

func init() {
	// Route => handler
	e.POST("/getStaticJson", jsonWriter)
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
