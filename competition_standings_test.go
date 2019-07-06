package footballapiv2

import (
	"os"
	"testing"
)

type StandingFilter struct {
	StandingType string
}

func TestCompetitionStandings(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := StandingFilter{
		StandingType: "HOME",
	}

	res, err := client.CompetitionStandings(2021, &filter)
	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
	if res.Standings[0].Type != filter.StandingType {
		t.Error()
	}
}

func TestCompetitionStandingsWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	res, err := client.CompetitionStandings(2021, nil)
	if err != nil {
		t.Error()
	}
	if res.Competition.ID != 2021 {
		t.Error()
	}
	if res.Standings[0].Type != "TOTAL" {
		t.Error()
	}
}
