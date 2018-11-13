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
	Id       int     `json:"id"`
	Title    string  `json:"title"`    // TODO: 削除予定
	Subtitle string  `json:"subtitle"` // TODO: 削除予定
	Cost     float64 `json:"cost"`
	Compare  float64 `json:"compare"`
	Rate     float64 `json:"rate"`
}

type DetailRequest struct {
	Message int `json:"message"`
}

func init() {
	// 検証用
	e.GET("/get", getRecords)
	e.GET("/set", setRecords)

	e.POST("/getReport", getReport)
	e.POST("/getDetail", getDetail)
	e.GET("/getJson", getDetail)

	rand.Seed(time.Now().UnixNano())
}

func randomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func getDetail(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var records []Report

	// Receive message
	message := new(DetailRequest)

	if reqErr := c.Bind(message); reqErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, reqErr.Error())
	}

	// TODO: 暫定ですから、とりあえず計算やってなくて、一日のデータを表示します
	timeDisplay := getRequestDay(message.Message)
	query := datastore.NewQuery("DataStore").Filter("dateTime >", timeDisplay).Limit(1)

	var results []DataStore
	if _, err := query.GetAll(ctx, &results); err != nil {

	}

	if len(results) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "No Record")
	}

	result := results[0]

	// 見積もり収益額
	reportRevenue := new(Report)
	reportRevenue.Id = 6
	reportRevenue.Title = "見積もり収益額"
	reportRevenue.Cost = float64(result.Revenue)
	records = append(records, *reportRevenue)

	// ページビュー
	reportPageView := new(Report)
	reportPageView.Id = 7
	reportPageView.Title = "ページビュー"
	reportPageView.Cost = float64(result.PageView)
	records = append(records, *reportPageView)

	// 表示回数
	reportViewCount := new(Report)
	reportViewCount.Id = 8
	reportViewCount.Title = "表示回数"
	reportViewCount.Cost = float64(result.ViewCount)
	records = append(records, *reportViewCount)

	// ページCTR
	reportClickCount := new(Report)
	reportClickCount.Id = 9
	reportClickCount.Title = "ページCTR"
	reportClickCount.Cost = result.CtRate
	records = append(records, *reportClickCount)

	// クリック率
	reportClickRate := new(Report)
	reportClickRate.Id = 10
	reportClickRate.Title = "クリック率"
	reportClickRate.Cost = result.ClickRate
	records = append(records, *reportClickRate)

	// カバレッジ
	reportCoverage := new(Report)
	reportCoverage.Id = 11
	reportCoverage.Title = "カバレッジ"
	reportCoverage.Cost = result.Coverage
	records = append(records, *reportCoverage)

	return c.JSON(http.StatusOK, records)
}

// !!!: 仮実装
func getRequestDay(index int) time.Time {
	return time.Now().AddDate(0, -index, 0)
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

	todayQuery := datastore.NewQuery("DataStore").Filter("dateTime >", timeYesterday).Filter("dateTime <=", timeNow)

	var results []DataStore
	if _, err := todayQuery.GetAll(ctx, &results); err != nil {

	}

	var costToday float64
	for i := range results {
		costToday = costToday + float64(results[i].Revenue)
	}

	reportToday.Cost = costToday
	records = append(records, *reportToday)

	// 昨日
	twoDaysAgo := timeNow.Add(-48 * time.Hour)

	reportYesterday := new(Report)
	reportYesterday.Id = 2
	reportYesterday.Title = "昨日"
	reportYesterday.Subtitle = "先週と同じ曜日との比較"

	yesterdayQuery := datastore.NewQuery("DataStore").Filter("dateTime >", twoDaysAgo).Filter("dateTime <=", timeYesterday)

	// TODO: リサイクル使用方法確認
	results = results[:0]
	if _, err := yesterdayQuery.GetAll(ctx, &results); err != nil {

	}

	var costYesterday float64
	for i := range results {
		costYesterday = costYesterday + float64(results[i].Revenue)
	}

	reportYesterday.Cost = costYesterday

	// TODO: その他の比較
	reportYesterday.Compare = 0.1
	reportYesterday.Rate = 0.1
	records = append(records, *reportYesterday)

	// 今月
	lastMonth := timeNow.AddDate(0, -1, 0)

	reportThisMonth := new(Report)
	reportThisMonth.Id = 3
	reportThisMonth.Title = "今月（現時点まで）"
	reportThisMonth.Subtitle = "先月の同じ日との比較"

	thisMonthQuery := datastore.NewQuery("DataStore").Filter("dateTime >", lastMonth).Filter("dateTime <=", timeNow)

	// TODO: リサイクル使用方法確認
	results = results[:0]
	if _, err := thisMonthQuery.GetAll(ctx, &results); err != nil {

	}

	var costThisMonth float64
	for i := range results {
		costThisMonth = costThisMonth + float64(results[i].Revenue)
	}

	reportThisMonth.Cost = costThisMonth
	records = append(records, *reportThisMonth)

	// 先月
	twoMonthsAgo := timeNow.AddDate(0, -2, 0)

	reportLastMonth := new(Report)
	reportLastMonth.Id = 4
	reportLastMonth.Title = "今月（現時点まで）"
	reportLastMonth.Subtitle = "先月の同じ日との比較"

	lastMonthQuery := datastore.NewQuery("DataStore").Filter("dateTime >", twoMonthsAgo).Filter("dateTime <=", lastMonth)

	// TODO: リサイクル使用方法確認
	results = results[:0]
	if _, err := lastMonthQuery.GetAll(ctx, &results); err != nil {

	}

	var costLastMonth float64
	for i := range results {
		costLastMonth = costLastMonth + float64(results[i].Revenue)
	}

	reportLastMonth.Cost = costLastMonth
	records = append(records, *reportLastMonth)

	// 全期間
	reportAll := new(Report)
	reportAll.Id = 5
	reportAll.Title = "全期間"

	allQuery := datastore.NewQuery("DataStore")

	// TODO: リサイクル使用方法確認
	results = results[:0]
	if _, err := allQuery.GetAll(ctx, &results); err != nil {

	}

	var costAll float64
	for i := range results {
		costAll = costAll + float64(results[i].Revenue)
	}

	reportAll.Cost = costAll
	records = append(records, *reportAll)

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
