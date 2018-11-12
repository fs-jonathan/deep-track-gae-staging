package app

import (
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type DataStore struct {
	id       int     `json:"id" datastore:"ID"`
	Title    string  `json:"title" datastore:"title,noindex"`
	Subtitle string  `json:"subtitle" datastore:"subtitle,noindex"`
	Cost     float64 `json:"cost" datastore:"cost,noindex"`
	Compare  float64 `json:"compare" datastore:"compare,noindex"`
	Rate     float64 `json:"rate" datastore:"rate,noindex"`
}

func init() {
	// 検証用
	e.GET("/get", getRecords)
	e.GET("/set", setRecords)
}

func getDefaultData() []DataStore {
	s0 := []DataStore{DataStore{1, "本日（現時点まで）", "", 9, 8, 7}}
	s1 := append(s0, DataStore{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
	s2 := append(s1, DataStore{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
	s3 := append(s2, DataStore{4, "先月", "先々月との比較", 93, 83, 72})
	records := append(s3, DataStore{5, "全期間", "", 20, 4, 434})

	return records
}

func getRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	// The query type
	query := datastore.NewQuery("Record")

	var results []DataStore
	if _, err := query.GetAll(ctx, &results); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func setRecords(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	records := getDefaultData()

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
