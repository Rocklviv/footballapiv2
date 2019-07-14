package footballapiv2

// Competitions response from data provider.
type Competitions struct {
	Competitions []Competition
}

// Competition list of competitions
type Competition struct {
	ID   int16
	Area struct {
		ID   int16
		Name string
	}
	Name          string
	Code          string `json:"code,omitempty"`
	EmblemURL     string `json:"emblemurl,omitempty"`
	Plan          string
	CurrentSeason struct {
		ID              int16
		StartDate       string
		EndDate         string
		CurrentMatchday int16
		Winner          struct{} `json:"winner,omitempty"`
	}
	NumberOfAvailableSeasons int16
	LastUpdated              string
}

// TeamsList of teams in league/competition
type TeamsList struct {
	Count       uint16
	Teams       []Team
	Filters     struct{}
	Competition Competition
	Season      Season
}

// Season information
type Season struct {
	ID              uint16
	StartDate       string
	EndDate         string
	CurrentMatchday uint
	AvailableStages []interface{}
	Winner          struct {
		ID        uint16
		Name      string
		ShortName string
		TLA       string
		CrestURL  string
	}
}

// Team infomration
type Team struct {
	ID   uint16
	Area struct {
		ID   uint16
		Name string
	}
	Name        string
	ShortName   string
	Tla         string
	CrestURL    string
	Address     string
	Phone       string
	Website     string
	Email       string
	Founded     uint16
	ClubColors  string
	Venue       string
	LastUpdated string
}

// ListStandings infomration
type ListStandings struct {
	Filters     struct{}
	Competition Competition
	Season      Season
	Standings   []Standing
}

// Standing datastruct
type Standing struct {
	Stage string
	Type  string
	Group string
	Table []Table
}

// Table - standing table of required competition
type Table struct {
	Position uint8
	Team     struct {
		ID       uint16
		Name     string
		CrestURL string
	}
	PlayedGames    uint8
	Won            uint8
	Draw           uint8
	Lost           uint8
	Points         int8
	GoalsFor       uint8
	GoalsAgaints   uint8
	GoalDifference uint8
}

// ListLeagueMatches represens matches in required league
type ListLeagueMatches struct {
	Count   uint8
	Filters struct {
		MatchDay string
	}
	Competition Competition
	Matches     []LeagueMatches
}

// LeagueMatches represents schema of league match
type LeagueMatches struct {
	ID          uint32
	Season      Season
	UtcData     string
	Status      string
	Matchday    uint8
	Stage       string
	Group       string
	LastUpdated string
	Score       Score
	HomeTeam    struct {
		ID    uint16
		Name  string
		Coach struct {
			ID             uint16
			Name           string
			CountryOfBirth string
			Nationality    string
		}
		LineUp []LineUp
		Bench  []Bench
	}
	AwayTeam struct {
		ID   uint16
		Name string
	}
	Referees []Referees
}

// Score match score
type Score struct {
	Winner   string
	Duration string
	FullTime struct {
		HomeTeam uint8
		AwayTeam uint8
	}
	HalfTime struct {
		HomeTeam uint8
		AwayTeam uint8
	}
	ExtraTime struct {
		HomeTeam uint8
		AwayTeam uint8
	}
	Penalties struct {
		HomeTeam uint8
		AwayTeam uint8
	}
}

// Referees represents referee
type Referees struct {
	ID          uint16
	Name        string
	Nationality string
}

// LineUp represents lineup of the match
type LineUp struct {
	ID          uint8
	Name        string
	Position    string
	ShirtNumber uint8
}

// Bench represents players on a bench
type Bench struct {
	ID          uint8
	Name        string
	Position    string
	ShirtNumber uint8
}

// ListLeagueScorers represents list of best scorers in league/competition
type ListLeagueScorers struct {
	Count   uint8
	Filters struct {
		Limit uint8
	}
	Competition Competition
	Season      Season
	Scorers     []Scorers
}

// Scorers represent exact goal scorer
type Scorers struct {
	Player Player
	Team   struct {
		ID   uint8
		Name string
	}
	NumberOfGoals uint16
}

// Player represents player information
type Player struct {
	ID             uint16
	Name           string
	FirtsName      string
	LastName       string
	DateOfBirth    string
	CountryOfBirth string
	Nationality    string
	Position       string
	ShirtNumber    uint8
	LastUpdated    string
}

// ListMatches represents matches for specified competition(s)
type ListMatches struct {
	Count   uint8
	Filters struct {
		DateFrom     string
		DateTo       string
		Permissions  string
		Competitions map[string]int
	}
	Matches []Match
}

// Match represents fixture/match for competition/matchday
type Match struct {
	ID          uint32
	Competition struct {
		ID   uint32
		Name string
	}
	Season      Season
	UtcDate     string
	Status      string
	Matchday    uint8
	Stage       string
	Group       string
	LastUpdated string
	Score       Score
	HomeTeam    struct {
		ID   uint32
		Name string
	}
	AwayTeam struct {
		ID   uint32
		Name string
	}
	Referees Referees
}

// Head2HeadMatch represents head-to-head match in required competition.
type Head2HeadMatch struct {
	Head2Head struct {
		NumberOfMatches uint8
		TotalGoals      uint16
		HomeTeam        struct {
			Wins   uint8
			Draws  uint8
			Losses uint8
		}
		AwayTeam struct {
			Wins   uint8
			Draws  uint8
			Losses uint8
		}
	}
	Match Match
}
