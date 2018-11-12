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

func init() {
	// 検証用
	e.GET("/get", getRecords)
	e.GET("/set", setRecords)

	rand.Seed(time.Now().UnixNano())
}

func randomFloat(min, max float64) float64 {
	return rand.Float64() * (max - min) + min
}

func getDefaultData() []DataStore {
	records := []DataStore{}
	date := time.Now()

	for i := 0; i < 20; i++ {
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
