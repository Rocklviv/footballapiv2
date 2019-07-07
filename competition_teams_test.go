package footballapiv2

import (
	"os"
	"strings"
	"testing"
)

type TeamsFilters struct {
	Season string
	Stage  string
}

func TestCompetitionTeamsWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.CompetitionTeams(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestCompetitionTeamsWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := TeamsFilters{
		Season: "2018",
	}

	res, err := client.CompetitionTeams(2021, &filter)
	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	} else if !strings.Contains(res.Season.StartDate, "2018") {
		t.Error()
	}
}

func TestCompetitionTeamsError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.CompetitionTeams(0, nil)

	if err == nil {
		t.Error()
	}
}

func TestCompetitionTeamsWStages(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := TeamsFilters{
		Stage: "GROUP_STAGE",
	}

	res, err := client.CompetitionTeams(2000, &filter)
	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2000 {
		t.Error()
	}
}
