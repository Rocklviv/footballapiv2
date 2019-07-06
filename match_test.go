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

func TestListMatchesWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.GetMatches(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestListMatchesWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := MatchesFilter{
		Matchday: "1",
		Season:   "2018",
	}

	res, err := client.GetMatches(2021, &filter)
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

func TestListMatchesError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.GetMatches(0, nil)

	if err == nil {
		t.Error()
	}
}
