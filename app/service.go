package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Record struct {
	Id       int     `json:"id" datastore:"ID"`
	Title    string  `json:"title" datastore:"title,noindex"`
	Subtitle string  `json:"subtitle" datastore:"subtitle,noindex"`
	Cost     float64 `json:"cost" datastore:"cost,noindex"`
	Compare  float64 `json:"compare" datastore:"compare,noindex"`
	Rate     float64 `json:"rate" datastore:"rate,noindex"`
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
	e.GET("/get", getRecords)
	e.GET("/set", setRecords)

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

func getRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	// The query type
	query := datastore.NewQuery("Record").Order("ID")

	var results []Record
	if _, err := query.GetAll(ctx, &results); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func setRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	records := getDefaultRecord()

	keys := []*datastore.Key{
		datastore.NewIncompleteKey(ctx, "Record", nil),
		datastore.NewIncompleteKey(ctx, "Record", nil),
		datastore.NewIncompleteKey(ctx, "Record", nil),
		datastore.NewIncompleteKey(ctx, "Record", nil),
		datastore.NewIncompleteKey(ctx, "Record", nil),
	}

	if _, err := datastore.PutMulti(ctx, keys, records); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, records)
}
