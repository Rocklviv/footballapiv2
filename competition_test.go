package footballapiv2

import (
	"os"
	"testing"
)

func TestGetCompetition(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	res, err := client.GetCompetition(2021)

	if err != nil {
		t.Error(err)
	}

	if res.ID != 2021 {
		t.Error()
	} else if res.Name != "Premier League" {
		t.Error()
	}
}

func TestGetCompetitionError(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	_, err := client.GetCompetition(0)

	if err == nil {
		t.Error()
	}
}
