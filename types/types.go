package types

//Game is the return data for a game
type Game struct {
	Home          string  `json:"HomeTeam"`
	Away          string  `json:"VisitorTeam"`
	HomeAbr       string  `json:"HomeTeamShort"`
	AwayAbr       string  `json:"VisitorTeamShort"`
	Time          string  `json:"EventTime"`
	OU            float32 `json:"OU"`
	HomePitcher   string  `json:"HomePrimaryPlayer"`
	AwayPItcher   string  `json:"VisitorPrimaryPlayer"`
	HomeHand      string  `json:"HomePitcherThrows"`
	AwayHand      string  `json:"VisitorPitcherThrows"`
	Stadium       string  `json:"StadiumName"`
	HRFactor      float64
	PPD           bool    `json:"IsPPD"`
	EventTitle    string  `json:"EventTitle"`
	EventSummary  string  `json:"EventSummary"`
	HomeRuns      float32 `json:"ProjHomeScore"`
	AwayRuns      float32 `json:"ProjVisitorScore"`
	IsFinal       bool
	WeatherGreen  bool    `json:"IsWeatherLevel1"`
	WeatherYellow bool    `json:"IsWeatherLevel2"`
	WeatherRed    bool    `json:"IsWeatherLevel3"`
	Weather       Weather `json:"EventWeather"`
	HomeBatStats  Matchup
	AwayBatStats  Matchup
}

//Weather is the return data for a games weather
type Weather struct {
	Summary     string  `json:"Summary"`
	RainPCT     float32 `json:"PrecipProbability"`
	Temp        float32 `json:"Temperature"`
	Humidity    float32 `json:"Humidity"`
	WindBearing int     `json:"AdjWindBearing"`
}

//GameJSON is the struct for storing json data
type GameJSON struct {
	Home          string  `json:"Home"`
	Away          string  `json:"Away"`
	HomeAbr       string  `json:"HomeAbr"`
	AwayAbr       string  `json:"AwayAbr"`
	Time          string  `json:"Time"`
	OU            float32 `json:"OU"`
	HomePitcher   string  `json:"HomePitcher"`
	AwayPItcher   string  `json:"AwayPitcher"`
	HomeHand      string  `json:"HomeHand"`
	AwayHand      string  `json:"AwayHand"`
	Stadium       string  `json:"Stadium"`
	HRFactor      float64 `json:"HRFactor"`
	PPD           bool    `json:"PPD"`
	EventTitle    string  `json:"EventTitle"`
	EventSummary  string  `json:"FinalScore"`
	HomeRuns      float32 `json:"HomeRuns"`
	AwayRuns      float32 `json:"AwayRuns"`
	IsFinal       bool    `json:"IsFinal"`
	WeatherGreen  bool    `json:"WeatherGreen"`
	WeatherYellow bool    `json:"WeatherYellow"`
	WeatherRed    bool    `json:"WeatherRed"`
	Weather       Weather `json:"Weather"`
	HomeBatStats  Matchup `json:"HomeBatStats"`
	AwayBatStats  Matchup `json:"AwayBatStats"`
}

//Matchup is the return struct for stattleship matchup data
type Matchup struct {
	BatK9            float64 `json:"BatterK9"`
	BatISO           float64 `json:"BatterISO"`
	AtBatsPerHomeRun float64 `json:"AtBatsPerHomeRun"`
}

//Results is the format for a document
type Results struct {
	Date  string     `json:"Date"`
	Slate []GameJSON `json:"Slate"`
}

//WeatherJSON is the return data for a games weather
type WeatherJSON struct {
	Summary     string  `json:"Summary"`
	RainPCT     float32 `json:"RainPCT"`
	Temp        float32 `json:"Temp"`
	Humidity    float32 `json:"Humidity"`
	WindBearing int     `json:"WindBearing"`
}

//TeamStatsBySeason is a struct for stattleship return data
type TeamStatsBySeason struct {
	Seasons []struct {
		ID                  string `json:"id"`
		CreatedAt           string `json:"created_at"`
		UpdatedAt           string `json:"updated_at"`
		Name                string `json:"name"`
		CurrentIntervalType string `json:"current_interval_type"`
		StartsOn            string `json:"starts_on"`
		EndsOn              string `json:"ends_on"`
		Slug                string `json:"slug"`
		LeagueID            string `json:"league_id"`
	} `json:"seasons"`
	Leagues []struct {
		ID               string      `json:"id"`
		CreatedAt        string      `json:"created_at"`
		UpdatedAt        string      `json:"updated_at"`
		Abbreviation     string      `json:"abbreviation"`
		Color            interface{} `json:"color"`
		MinutesPerPeriod interface{} `json:"minutes_per_period"`
		Name             string      `json:"name"`
		Periods          interface{} `json:"periods"`
		Slug             string      `json:"slug"`
		Sport            string      `json:"sport"`
	} `json:"leagues"`
	Teams []struct {
		ID         string   `json:"id"`
		CreatedAt  string   `json:"created_at"`
		UpdatedAt  string   `json:"updated_at"`
		Color      string   `json:"color"`
		Colors     []string `json:"colors"`
		Hashtag    string   `json:"hashtag"`
		Hashtags   []string `json:"hashtags"`
		Location   string   `json:"location"`
		Name       string   `json:"name"`
		Nickname   string   `json:"nickname"`
		Latitude   float64  `json:"latitude"`
		Longitude  float64  `json:"longitude"`
		Slug       string   `json:"slug"`
		DivisionID string   `json:"division_id"`
		LeagueID   string   `json:"league_id"`
	} `json:"teams"`
	Divisions []struct {
		ID           string `json:"id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		Name         string `json:"name"`
		ConferenceID string `json:"conference_id"`
	} `json:"divisions"`
	Conferences []struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Name      string `json:"name"`
		LeagueID  string `json:"league_id"`
	} `json:"conferences"`
	TeamSeasonStats []struct {
		ID                               string  `json:"id"`
		CreatedAt                        string  `json:"created_at"`
		UpdatedAt                        string  `json:"updated_at"`
		IntervalType                     string  `json:"interval_type"`
		StatisticsOn                     string  `json:"statistics_on"`
		Assists                          float64 `json:"assists"`
		WildPitchesWhileCatching         float64 `json:"wild_pitches_while_catching"`
		DoublePlays                      float64 `json:"double_plays"`
		FieldingErrors                   float64 `json:"fielding_errors"`
		FieldingPct                      float64 `json:"fielding_pct"`
		PassedBalls                      float64 `json:"passed_balls"`
		Putouts                          float64 `json:"putouts"`
		TotalChances                     int     `json:"total_chances"`
		TriplePlays                      int     `json:"triple_plays"`
		AtBats                           float64 `json:"at_bats"`
		AtBatsRisp                       float64 `json:"at_bats_risp"`
		AtBatsPerHomeRun                 float64 `json:"at_bats_per_home_run"`
		AtBatsPerStrikeOut               float64 `json:"at_bats_per_strike_out"`
		PlateAppearances                 float64 `json:"plate_appearances"`
		BattingAverage                   float64 `json:"batting_average"`
		Babip                            float64 `json:"babip"`
		WalksPerStrikeOut                float64 `json:"walks_per_strike_out"`
		WalksPerPlateAppearance          float64 `json:"walks_per_plate_appearance"`
		BallsInPlay                      float64 `json:"balls_in_play"`
		Gofo                             float64 `json:"gofo"`
		HitsRisp                         int     `json:"hits_risp"`
		Iso                              float64 `json:"iso"`
		RunnersLeftOnBase                float64 `json:"runners_left_on_base"`
		LobRisp2Out                      float64 `json:"lob_risp_2out"`
		OnBasePercentage                 float64 `json:"on_base_percentage"`
		Walks                            int     `json:"walks"`
		Doubles                          int     `json:"doubles"`
		FieldersChoice                   int     `json:"fielders_choice"`
		Hits                             int     `json:"hits"`
		HitByPitches                     int     `json:"hit_by_pitches"`
		HomeRuns                         int     `json:"home_runs"`
		IntentionalWalks                 int     `json:"intentional_walks"`
		ReachedOnError                   int     `json:"reached_on_error"`
		Singles                          int     `json:"singles"`
		Triples                          int     `json:"triples"`
		TotalBases                       int     `json:"total_bases"`
		Ops                              float64 `json:"ops"`
		Balls                            int     `json:"balls"`
		Dirtballs                        int     `json:"dirtballs"`
		FoulBalls                        int     `json:"foul_balls"`
		IntentionalBalls                 int     `json:"intentional_balls"`
		StrikesLooking                   int     `json:"strikes_looking"`
		StrikesSwinging                  int     `json:"strikes_swinging"`
		Strikes                          int     `json:"strikes"`
		FlyOutsIntoDoublePlays           int     `json:"fly_outs_into_double_plays"`
		FlyOuts                          int     `json:"fly_outs"`
		GroundOutsIntoDoublePlays        int     `json:"ground_outs_into_double_plays"`
		GroundOuts                       int     `json:"ground_outs"`
		StrikeoutsLooking                int     `json:"strikeouts_looking"`
		StrikeoutsSwinging               int     `json:"strikeouts_swinging"`
		Strikeouts                       int     `json:"strikeouts"`
		LineOutsIntoDoublePlays          int     `json:"line_outs_into_double_plays"`
		LineOuts                         int     `json:"line_outs"`
		PopOuts                          int     `json:"pop_outs"`
		SacrificeFlies                   int     `json:"sacrifice_flies"`
		SacrificeHits                    int     `json:"sacrifice_hits"`
		PitchesFaced                     float64 `json:"pitches_faced"`
		Rbi                              float64 `json:"rbi"`
		EarnedRuns                       int     `json:"earned_runs"`
		Runs                             int     `json:"runs"`
		UnearnedRuns                     int     `json:"unearned_runs"`
		SecondaryAverage                 float64 `json:"secondary_average"`
		SluggingPct                      float64 `json:"slugging_pct"`
		CaughtStealing                   int     `json:"caught_stealing"`
		StolenBasePct                    float64 `json:"stolen_base_pct"`
		StolenBases                      int     `json:"stolen_bases"`
		TeamLeftOnBase                   int     `json:"team_left_on_base"`
		ExtraBaseHits                    float64 `json:"extra_base_hits"`
		BattersFaced                     float64 `json:"batters_faced"`
		Era                              float64 `json:"era"`
		PitchingErrors                   float64 `json:"pitching_errors"`
		BlownSave                        int     `json:"blown_save"`
		CompleteGames                    int     `json:"complete_games"`
		Holds                            int     `json:"holds"`
		Losses                           int     `json:"losses"`
		QualityStarts                    int     `json:"quality_starts"`
		Saves                            int     `json:"saves"`
		Shutouts                         int     `json:"shutouts"`
		SaveOpportunities                int     `json:"save_opportunities"`
		TeamShutouts                     int     `json:"team_shutouts"`
		Wins                             int     `json:"wins"`
		GroundBallToFlyBallRatio         float64 `json:"ground_ball_to_fly_ball_ratio"`
		InningsPitchedTotalOuts          int     `json:"innings_pitched_total_outs"`
		InningsPitched                   int     `json:"innings_pitched"`
		StrikeoutsPer9Innings            float64 `json:"strikeouts_per_9_innings"`
		StrikeoutsPerWalk                float64 `json:"strikeouts_per_walk"`
		PitcherRunnersLeftOnBase         float64 `json:"pitcher_runners_left_on_base"`
		OpponentsBattingAverage          float64 `json:"opponents_batting_average"`
		WalksAllowed                     int     `json:"walks_allowed"`
		DoublesAllowed                   int     `json:"doubles_allowed"`
		FieldersChoiceAllowed            int     `json:"fielders_choice_allowed"`
		HitsAllowed                      int     `json:"hits_allowed"`
		HitBatters                       int     `json:"hit_batters"`
		HomeRunsAllowed                  int     `json:"home_runs_allowed"`
		IntentionalWalksAllowed          int     `json:"intentional_walks_allowed"`
		ReachedOnErrorAllowed            int     `json:"reached_on_error_allowed"`
		SinglesAllowed                   int     `json:"singles_allowed"`
		TriplesAllowed                   int     `json:"triples_allowed"`
		TotalBasesAllowed                int     `json:"total_bases_allowed"`
		BallsThrowns                     int     `json:"balls_throwns"`
		DirtballsThrowns                 int     `json:"dirtballs_throwns"`
		FoulBallsThrowns                 int     `json:"foul_balls_throwns"`
		IntentionalBallsThrowns          int     `json:"intentional_balls_throwns"`
		StrikesThrownLooking             int     `json:"strikes_thrown_looking"`
		StrikesThrownSwinging            int     `json:"strikes_thrown_swinging"`
		StrikesThrown                    int     `json:"strikes_thrown"`
		FlyOutsIntoDoublePlaysAllowed    int     `json:"fly_outs_into_double_plays_allowed"`
		FlyOutsAllowed                   int     `json:"fly_outs_allowed"`
		GroundOutsIntoDoublePlaysAllowed int     `json:"ground_outs_into_double_plays_allowed"`
		GroundOutsAllowed                int     `json:"ground_outs_allowed"`
		StrikeoutsLookingAllowed         int     `json:"strikeouts_looking_allowed"`
		StrikeoutsSwingingAllowed        int     `json:"strikeouts_swinging_allowed"`
		StrikeoutsAllowed                int     `json:"strikeouts_allowed"`
		LineOutsIntoDoublePlaysAllowed   int     `json:"line_outs_into_double_plays_allowed"`
		LineOutsAllowed                  int     `json:"line_outs_allowed"`
		PopOutsAllowed                   int     `json:"pop_outs_allowed"`
		SacrificeFliesAllowed            int     `json:"sacrifice_flies_allowed"`
		SacrificeHitsAllowed             int     `json:"sacrifice_hits_allowed"`
		PitchesThrown                    float64 `json:"pitches_thrown"`
		EarnedRunsAllowed                int     `json:"earned_runs_allowed"`
		RunsAllowed                      int     `json:"runs_allowed"`
		UnearnedRunsAllowed              int     `json:"unearned_runs_allowed"`
		RunnersCaughtStealing            int     `json:"runners_caught_stealing"`
		StolenBasesAllowed               int     `json:"stolen_bases_allowed"`
		Whip                             float64 `json:"whip"`
		WildPitches                      float64 `json:"wild_pitches"`
		SeasonID                         string  `json:"season_id"`
		TeamID                           string  `json:"team_id"`
	} `json:"team_season_stats"`
}

//Roster is the return data from a return search
type Roster struct {
	Team    string `json:"Team"`
	Players []struct {
		FirstName            string `json:"first_name"`
		LastName             string `json:"last_name"`
		Handedness           string `json:"handedness"`
		Name                 string `json:"name"`
		PositionAbbreviation string `json:"position_abbreviation"`
		Slug                 string `json:"slug"`
	} `json:"players"`
}

//MgoRoster is type for passing to mgodb
type MgoRoster struct {
	Team    string `json:"Team"`
	Players []struct {
		FirstName            string `json:"FirstName"`
		LastName             string `json:"LastName"`
		Handedness           string `json:"Handedness"`
		Name                 string `json:"Name"`
		PositionAbbreviation string `json:"PositionAbbreviation"`
		Slug                 string `json:"Slug"`
	} `json:"Players"`
}
