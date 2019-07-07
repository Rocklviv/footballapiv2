package footballapiv2

import (
	"os"
	"testing"
)

type Filter struct {
	Areas string
	Plan  string
}

func TestWithoutFilter(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	res, err := client.GetCompetitions(nil)
	if err != nil {
		t.Error()
	}
	if len(res) <= 1 {
		t.Error()
	}
}

func TestFilterTier(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := Filter{
		Plan: "TIER_ONE",
	}
	res, err := client.GetCompetitions(&filter)
	if err != nil {
		t.Error()
	}
	if len(res) <= 11 {
		t.Error()
	}
}

func TestFilterArea(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := Filter{
		Areas: "2077",
	}

	res, err := client.GetCompetitions(&filter)
	if err != nil {
		t.Error()
	}
	if len(res) <= 1 {
		t.Error()
	}
}

func TestWithBothFilters(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := Filter{
		Areas: "2077,2061",
		Plan:  "TIER_ONE",
	}
	res, err := client.GetCompetitions(&filter)
	if err != nil {
		t.Error()
	}
	if len(res) <= 1 {
		t.Error()
	}
}
