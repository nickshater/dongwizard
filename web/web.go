//Package web manages query handlers
package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	data "github.com/nickshater/dongwizard/data"
	"github.com/nickshater/dongwizard/db"
)

//TodayHandler returns the JSON data at the index
func TodayHandler(w http.ResponseWriter, r *http.Request) {
	date := data.GetTodaysDate()
	d := data.GetSlateData(date)
	w.Write(d)
}

//TomorrowHandler returns the json data at the /tomorrow endpoint
func TomorrowHandler(w http.ResponseWriter, r *http.Request) {
	date := data.GetTomorrowsDate()
	d := data.GetSlateData(date)
	w.Write(d)
}

//YesterdayHandler returns the json data at the /yesterday endpoint
func YesterdayHandler(w http.ResponseWriter, r *http.Request) {
	date := data.GetYesterdaysDate()
	d := data.GetSlateData(date)
	w.Write(d)
}

//ByDateHandler returns the json data for /[date] endpoint
func ByDateHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()
	date := u.Get("d")

	d := data.GetSlateData(date)
	w.Write(d)
}

//ArchiveByDateHandler calls to mongodb to return a result for a specific date
func ArchiveByDateHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()
	date := u.Get("d")
	res := data.GetDateDataWrite(date)
	mn := db.ArchiveQueryByDate(res)
	j, err := json.MarshalIndent(mn, "", "    ")
	if err != nil {
		fmt.Println("Marshal from mongo error", err)
	}
	w.Write(j)
}

//AddByDateHandler adds a document to the slate db by date
func AddByDateHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()
	date := u.Get("d")
	data := data.GetDateDataWrite(date)
	result := db.UpdateMngByDate(data)
	j, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Marshal from upsert error", err)
	}
	w.Write(j)
}
