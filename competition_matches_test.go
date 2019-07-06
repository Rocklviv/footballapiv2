package footballapiv2

import (
	"os"
	"strings"
	"testing"
)

type MatchesFilter struct {
	Matchday string
	DateFrom string
	DateTo   string
	Stage    string
	Status   string
	Group    string
	Season   string
}

func TestCompetitionMatchesWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.CompetitionMatches(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestCompetitionMatchesWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := MatchesFilter{
		Matchday: "1",
		Season:   "2018",
	}

	res, err := client.CompetitionMatches(2021, &filter)
	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
	if res.Filters.MatchDay != filter.Matchday {
		t.Error()
	}
	if !strings.Contains(res.Matches[0].Season.StartDate, filter.Season) {
		t.Error()
	}
}

func TestCompetitionMatchesError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.CompetitionMatches(0, nil)

	if err == nil {
		t.Error()
	}
}
