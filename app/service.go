package app

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type DataStore struct {
	RecordDate time.Time `datastore:"dateTime"`
	Revenue    int       `datastore:"revenue"`
	PageView   int       `datastore:"pageView"`
	ViewCount  int       `datastore:"viewCount"`
	CtRate     float64   `datastore:"ctRate"`
	ClickRate  float64   `datastore:"clickRate"`
	Coverage   float64   `datastore:"coverage"`
}

type Report struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Subtitle   string    `json:"subtitle"`
	Cost       int       `json:"cost"`
	Compare    float64   `json:"compare"`
	Rate       float64   `json:"rate"`
	RecordDate time.Time `json:"recordDate"`
}

func init() {
	// 検証用
	e.GET("/get", getRecords)
	e.GET("/set", setRecords)

	e.POST("/getJson", getReport)
	e.GET("/getReport", getReport)

	rand.Seed(time.Now().UnixNano())
}

func randomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func getReport(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var records []Report

	// TODO: 日時範囲選択は後ほど対応、暫定1ヶ月のデータを利用します
	// !!!: とりあえず日本語のみを想定します

	// 本日
	timeNow := time.Now()
	timeYesterday := timeNow.Add(-24 * time.Hour)

	reportToday := new(Report)
	reportToday.Id = 1
	reportToday.Title = "本日"
	reportToday.Subtitle = "（現時点まで）"
	reportToday.RecordDate = timeNow

	todayQuery := datastore.NewQuery("DataStore").Filter("dateTime >", timeYesterday).Filter("dateTime <=", timeNow)

	var results []DataStore
	if _, err := todayQuery.GetAll(ctx, &results); err != nil {

	}

	var costToday int
	for i := range results {
		costToday = costToday + results[i].Revenue
	}

	reportToday.Cost = costToday
	records = append(records, *reportToday)

	return c.JSON(http.StatusOK, records)
}

func getDefaultData() []DataStore {
	records := []DataStore{}
	date := time.Now()

	for i := 0; i < 31; i++ {
		n := DataStore{date, rand.Intn(100), rand.Intn(200), rand.Intn(400), randomFloat(0, 1), randomFloat(0, 1), randomFloat(0, 1)}
		records = append(records, n)

		date = date.AddDate(0, 0, -1)
	}

	return records
}

func getRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	// The query type
	query := datastore.NewQuery("DataStore")

	var results []DataStore
	if _, err := query.GetAll(ctx, &results); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func setRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	records := getDefaultData()
	recordLength := len(records)

	keys := make([]*datastore.Key, recordLength)

	for i := range keys {
		keys[i] = datastore.NewIncompleteKey(ctx, "DataStore", nil)
	}

	if _, err := datastore.PutMulti(ctx, keys, records); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, records)
}
