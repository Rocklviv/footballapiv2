package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// CompetitionMatches returns list of matches from required competition
func (c *Client) CompetitionMatches(competitionID int16, values interface{}) (*ListLeagueMatches, error) {
	var listLeagueMatches ListLeagueMatches
	var res *json.Decoder
	var err error
	path := fmt.Sprintf("competitions/%d/matches", competitionID)

	if values == nil {
		res, err = c.doRequest("GET", path, nil)
	} else {
		res, err = c.doRequest("GET", path, structToMap(values))
	}

	if err != nil {
		return nil, err
	}

	err = res.Decode(&listLeagueMatches)
	if err != nil {
		return nil, err
	}
	return &listLeagueMatches, nil
}
