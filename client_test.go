package footballapiv2

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	apikey := os.Getenv("FOOTBALL_API_KEY")
	client := NewClient(apikey)

	if client.apikey != apikey {
		t.Error()
	} else if client.apikey == "" {
		t.Error()
	}
}

func TestNewClientWithMaxRPM(t *testing.T) {
	apikey := os.Getenv("FOOTBALL_API_KEY")
	client := NewClient(apikey).SetMaxRPM(30)

	if client.rpm != 30 {
		t.Error()
	}
}

func TestRequest(t *testing.T) {
	var v Competitions
	apikey := os.Getenv("FOOTBALL_API_KEY")
	client := NewClient(apikey)

	res, err := client.doRequest("GET", "competitions", nil)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error()
	}
	res.Decode(&v)
}
