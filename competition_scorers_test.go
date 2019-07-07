package footballapiv2

import (
	"os"
	"strings"
	"testing"
)

type ScorersFilter struct {
	Season string
	Limit  uint8
}

func TestCompetitionScorersWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := ScorersFilter{
		Season: "2018",
		Limit:  20,
	}

	res, err := client.CompetitionScorers(2021, &filter)
	if err != nil {
		t.Error()
	}
	if !strings.Contains(res.Season.StartDate, filter.Season) {
		t.Error()
	}
	if res.Competition.ID != 2021 {
		t.Error()
	}
	if res.Filters.Limit != filter.Limit {
		t.Error()
	}
}

func TestCompetitionScorersWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.CompetitionScorers(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestCompetitionScorersError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.CompetitionScorers(0, nil)

	if err == nil {
		t.Error()
	}
}
