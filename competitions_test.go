package footballapiv2

import (
	"os"
	"testing"
)

type Filter struct {
	Areas []int
	Plan  string
}

func TestWithoutFilter(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))

	res := client.GetCompetitions(nil)
	if len(res) <= 1 {
		t.Error()
	}
}

func TestFilterTier(t *testing.T) {
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	filter := Filter{
		Plan: "TIER_ONE",
	}
	res := client.GetCompetitions(&filter)
	if len(res) <= 11 {
		t.Error()
	}
}

func TestFilterArea(t *testing.T) {
	var areas []int
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	areas = append(areas, 2077)
	filter := Filter{
		Areas: areas,
	}

	res := client.GetCompetitions(&filter)
	if len(res) <= 1 {
		t.Error()
	}
}

func TestWithBothFilters(t *testing.T) {
	var areas []int
	client := NewClient(os.Getenv("FOOTBALL_API_KEY"))
	areas = append(areas, 2077)
	filter := Filter{
		Areas: areas,
		Plan:  "TIER_ONE",
	}
	t.Log(filter)
	res := client.GetCompetitions(&filter)
	if len(res) <= 1 {
		t.Error()
	}
	t.Log(len(res))
	t.Log(res)
}
