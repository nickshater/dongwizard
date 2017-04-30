package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nickshater/ns/web"
)

func main() {
	// data.GetPitcherData()
	http.HandleFunc("/", web.TodayHandler)
	http.HandleFunc("/tomorrow", web.TomorrowHandler)
	http.HandleFunc("/yesterday", web.YesterdayHandler)
	http.HandleFunc("/date", web.ByDateHandler)
	http.HandleFunc("/addDate", web.AddByDateHandler)
	http.HandleFunc("/archive", web.ArchiveByDateHandler)
	fmt.Println("Running Server on port :3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
