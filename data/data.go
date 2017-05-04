//Package data collects html data and saves it to a file
package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/nickshater/dongwizard/db"
	"github.com/nickshater/dongwizard/types"
)

const (
	seasonStart = "4_2_2017"
	seasonEnd   = "10_1_2017"
)

//GetTodaysDate returns todays date
func GetTodaysDate() string {
	time := time.Now().Local()
	t := time.Format("01_02_2006")
	return t
}

//GetTomorrowsDate returns tomorrows date
func GetTomorrowsDate() string {
	time := time.Now().Local()
	tom := time.AddDate(0, 0, 1)
	t := tom.Format("01_02_2006")
	return t
}

//GetYesterdaysDate returns yesterdays date
func GetYesterdaysDate() string {
	time := time.Now().Local()
	tom := time.AddDate(0, 0, -1)
	t := tom.Format("01_02_2006")
	return t
}

//GetSlateData returns html info about a slate of MLB games
func GetSlateData(date string) []byte {
	url := "http://www.fantasylabs.com/api/sportevents/3/" + date
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read error", err)
	}
	var d []types.Game
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	s := MergeSlateData(d)

	var r types.Results

	r.Date = date
	r.Slate = s

	j, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Println("Marshal error", err)
	}
	return j
}

//GetDateDataWrite returns an object for writing document to mongoDB
func GetDateDataWrite(date string) types.Results {
	url := "http://www.fantasylabs.com/api/sportevents/3/" + date
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read error", err)
	}
	var d []types.Game
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	s := MergeSlateData(d)

	var r types.Results

	r.Date = date
	r.Slate = s
	return r
}

//MergeSlateData converts the seperate results into a GameJSON struct for return
func MergeSlateData(d []types.Game) []types.GameJSON {
	var s []types.GameJSON
	for i := 0; i < len(d); i++ {
		re := regexp.MustCompile("FINAL")
		if d[i].AwayHand == "" {
			d[i].AwayHand = "(R)"
		}
		if d[i].HomeHand == "" {
			d[i].HomeHand = "(R)"
		}
		isFinal := re.FindString(d[i].EventSummary)
		if isFinal != "" {
			d[i].IsFinal = true
		}

		lh := strings.ToLower(d[i].HomeAbr)
		la := strings.ToLower(d[i].AwayAbr)
		if lh == "wsh" {
			lh = "was"
		}
		if la == "wsh" {
			la = "was"
		}
		if lh == "cws" {
			lh = "chw"
		}
		if la == "cws" {
			la = "chw"
		}
		if lh == "lad" {
			lh = "la"
		}
		if la == "lad" {
			la = "la"
		}

		hLstName := strings.Fields(d[i].HomePitcher)
		// aLstName := strings.Fields(d[i].AwayPitcher)
		hidx := (len(hLstName) - 1)
		htmslg := "mlb-" + lh
		fmt.Println(hLstName[hidx])
		// aidx := (len(aLstName) - 1)
		hpslug := db.GetPitcherSlug(htmslg, hLstName[hidx])
		// apslug := db.GetPitcherSlug("mlb-"+la, aLstName[len(aLstName)-1])
		d[i].HomeSlug = hpslug
		// d[i].AwaySlug = apslug
		hm := GetTeamSeasonData(lh)
		d[i].HomeBatStats.BatK9 = hm.TeamSeasonStats[0].StrikeoutsPer9Innings
		d[i].HomeBatStats.BatISO = hm.TeamSeasonStats[0].Iso
		d[i].HomeBatStats.AtBatsPerHomeRun = hm.TeamSeasonStats[0].AtBatsPerHomeRun

		am := GetTeamSeasonData(la)
		d[i].AwayBatStats.BatK9 = am.TeamSeasonStats[0].StrikeoutsPer9Innings
		d[i].AwayBatStats.BatISO = am.TeamSeasonStats[0].Iso
		d[i].AwayBatStats.AtBatsPerHomeRun = am.TeamSeasonStats[0].AtBatsPerHomeRun

		hrf := db.GetParkFactorFromDB(lh)

		d[i].HRFactor = hrf

		s = append(s, types.GameJSON(d[i]))
	}
	return s
}

//GetTeamSeasonData connects to the stattleship API to return k/9 and calc hr/9
func GetTeamSeasonData(team string) types.TeamStatsBySeason {
	var uri = "https://api.stattleship.com/baseball/mlb/team_season_stats?on=today&team_id=mlb-" + team
	client := &http.Client{}
	key := os.Getenv("STATTLESHIP")

	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", key)
	request.Header.Set("Accept", "application/vnd.stattleship.com; version=1")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("GetTeamSeason error ", err)
	}
	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetTeamSeasonData Body error ", err)
	}
	defer resp.Body.Close()
	var m types.TeamStatsBySeason
	err = json.Unmarshal(results, &m)
	if err != nil {
		fmt.Println("Matchup Unmarshal error", err)
	}
	return m
}

//PrintTheResults is used to return sample json from an api
func PrintTheResults() {
	var uri = "https://api.stattleship.com/baseball/mlb/rosters?team_id=mlb-tor"

	client := &http.Client{}
	key := os.Getenv("STATTLESHIP")

	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", key)
	request.Header.Set("Accept", "application/vnd.stattleship.com; version=1")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("GetTeamSeason error ", err)
	}
	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetTeamSeasonData Body error ", err)
	}
	defer resp.Body.Close()

	fmt.Println(string(results))
}

//GetPitcherData gets pitcher slug and name for searching for stats
//Need to do this concurrently in the future but since this is a one time function run it
//Can go on the backburner
func GetPitcherData() {
	teams := db.ReturnTeamsForRosterSearch()
	var r []types.MgoRoster

	for i := 0; i < len(teams); i++ {
		var d types.Roster
		var uri = "https://api.stattleship.com/baseball/mlb/rosters?team_id=" + teams[i]
		client := &http.Client{}
		key := os.Getenv("STATTLESHIP")

		d.Team = teams[i]
		request, err := http.NewRequest("GET", uri, nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", key)
		request.Header.Set("Accept", "application/vnd.stattleship.com; version=1")
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("GetTeamSeason error ", err)
		}
		results, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("GetTeamSeasonData Body error ", err)
		}
		defer resp.Body.Close()

		err = json.Unmarshal(results, &d)
		if err != nil {
			fmt.Println("Unmarshal error", err)
		}
		r = append(r, types.MgoRoster(d))
	}
	db.UpdateMngRosters(r)
}
