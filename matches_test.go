package footballapiv2

import (
	"os"
	"testing"
)

type DateFilter struct {
	DateFrom string
}

type NormalFilter struct {
	DateFrom     string
	DateTo       string
	Competitions uint16
}

type MatchesCompetitionFilter struct {
	Competitions uint16
}

func TestListMatchesWOFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	_, err := client.ListOfMatches(nil)

	if err != nil {
		t.Error()
	}
}

func TestListMatchesWFiltersError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := DateFilter{
		DateFrom: "2018-08-23",
	}

	_, err := client.ListOfMatches(&filter)

	if err == nil {
		t.Error()
	}
	t.Log(err)
}

func TestListMatchesWFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := NormalFilter{
		DateFrom:     "2018-08-23",
		DateTo:       "2018-08-23",
		Competitions: 2013,
	}

	res, err := client.ListOfMatches(&filter)
	if err != nil {
		t.Error()
	}

	if res.Filters.DateFrom != filter.DateFrom {
		t.Error()
	}

	if res.Filters.DateTo != filter.DateTo {
		t.Error()
	}

	if res.Matches[0].Competition.ID != filter.Competitions {
		t.Error()
	}
}

func TestListMatchesWCompetitionFilter(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	filter := NormalFilter{
		Competitions: 2013,
	}

	res, err := client.ListOfMatches(&filter)
	if err != nil {
		t.Error(err)
	}

	if res.Count < 0 {
		t.Error()
	}
}
