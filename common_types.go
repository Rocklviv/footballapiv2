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
	Winner          string
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
	Founded     uint8
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
		ID       uint8
		Name     string
		CrestURL string
	}
	PlayedGames    uint8
	Won            uint8
	Draw           uint8
	Lost           uint8
	Points         uint8
	GoalsFor       uint16
	GoalsAgaints   uint16
	GoalDifference uint16
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
	ID          uint16
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
