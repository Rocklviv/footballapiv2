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
	Season      struct {
		ID              uint16
		StartDate       string
		EndDate         string
		CurrentMatchday uint
		AvailableStages []interface{}
		Winner          string
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
	Founded     uint8
	ClubColors  string
	Venue       string
	LastUpdated string
}
