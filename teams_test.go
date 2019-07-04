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

func TestGetTeamsWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.GetTeams(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestGetTeamsWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := TeamsFilters{
		Season: "2018",
	}

	res, err := client.GetTeams(2021, &filter)
	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	} else if !strings.Contains(res.Season.StartDate, "2018") {
		t.Error()
	}
}

func TestGetTeamsError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.GetTeams(0, nil)

	if err == nil {
		t.Error()
	}
}
