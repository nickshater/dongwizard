package db

import (
	"fmt"

	"github.com/nickshater/dongwizard/etcd"
	"github.com/nickshater/dongwizard/types"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetSession creates a mongoDB session
func GetSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("mongo session error", err)
	}
	user := etcd.GetValue("mongousername")
	pwd := etcd.GetValue("mongopassword")
	s.DB("dongwizard").Login(user, pwd)
	return s
}

//ArchiveQueryByDate returns results from a date query
func ArchiveQueryByDate(r types.Results) types.Results {
	s := GetSession()
	defer s.Close()

	result := types.Results{}
	c := s.DB("dongwizard").C("slates")
	err := c.Find(bson.M{"date": r.Date}).One(&result)

	if err != nil {
		fmt.Println("ArchiveByDateInsertError", err)
	}

	return result
}

//UpdateMngByDate will add a result to the db if not existing and update if it does
func UpdateMngByDate(results types.Results) types.Results {
	s := GetSession()
	defer s.Close()

	c := s.DB("dongwizard").C("slates")
	_, err := c.Upsert(bson.M{"date": results.Date}, &results)
	if err != nil {
		fmt.Println("UpdateMngByDate Error", err)
	}
	return results
}

//GetParkFactorFromDB returns the park factor for the home teams stadium
func GetParkFactorFromDB(team string) float64 {
	s := GetSession()
	defer s.Close()

	c := s.DB("dongwizard").C("homerunFactor")
	var result struct {
		Park     string  `bson:"park" json:"park"`
		HomeTeam string  `bson:"home_team" json:"home_team"`
		Factor   float64 `bson:"hr_factor" json:"hr_factor"`
	}

	err := c.Find(bson.M{"home_team": team}).Select(bson.M{"hr_factor": 1}).One(&result)
	if err != nil {
		fmt.Println("Park Factor Error ", err)
	}
	return result.Factor
}

//ReturnTeamsForRosterSearch creates a slice of team abrs for roster querying
func ReturnTeamsForRosterSearch() []string {
	s := GetSession()
	defer s.Close()

	c := s.DB("dongwizard").C("homerunFactor")
	var res []struct {
		Park     string  `bson:"park" json:"park"`
		HomeTeam string  `bson:"home_team" json:"home_team"`
		Factor   float64 `bson:"hr_factor" json:"hr_factor"`
	}

	err := c.Find(nil).All(&res)
	if err != nil {
		fmt.Println("Park Factor Error ", err)
	}
	var teams []string
	for a := range res {
		teams = append(teams, "mlb-"+res[a].HomeTeam)
	}
	return teams
}

//UpdateMngRosters will add a roster to the db if not existing and update if it does
func UpdateMngRosters(rosters []types.MgoRoster) {
	s := GetSession()
	defer s.Close()

	c := s.DB("dongwizard").C("rosters")
	for i := 0; i < len(rosters); i++ {
		_, err := c.Upsert(bson.M{"Team": rosters[i].Team}, &rosters[i])
		if err != nil {
			fmt.Println("UpdateMngByDate Error", err)
		}
	}
}

//GetPitcherSlug returns the player slug for a searched for pitcher
func GetPitcherSlug(team string, lastname string) string {
	s := GetSession()
	defer s.Close()
	var r struct {
		Players []struct {
			Firstname            string `bson:"firstname"`
			Lastname             string `bson:"lastname"`
			Handedness           string `bson:"handedness"`
			Name                 string `bson:"name"`
			Positionabbreviation string `bson:"positionabbreviation"`
			Slug                 string `bson:"slug"`
		} `bson:"players"`
	}
	c := s.DB("dongwizard").C("rosters")
	err := c.Find(bson.M{"team": team}).Select(bson.M{"players": bson.M{"$elemMatch": bson.M{"lastname": lastname}}}).One(&r)

	var slug string
	if err != nil {
		fmt.Println("GetPitcherSlug error ", err)
	}
	if len(r.Players) == 0 {
		slug = "unknown"
	} else {
		slug = r.Players[0].Slug
	}
	return slug
}
