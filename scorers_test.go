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

func TestGetListOfScorersWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := ScorersFilter{
		Season: "2018",
		Limit:  20,
	}

	res, err := client.GetListOfScorers(2021, &filter)
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

func TestGetListOfScorersWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.GetListOfScorers(2021, nil)

	if err != nil {
		t.Error()
	}

	if res.Competition.ID != 2021 {
		t.Error()
	}
}

func TestGetListOfScorersError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.GetListOfScorers(0, nil)

	if err == nil {
		t.Error()
	}
}
