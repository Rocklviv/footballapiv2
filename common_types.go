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
