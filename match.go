package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// GetMatches returns list of matches from required competition
func (c *Client) GetMatches(competitionID int, values interface{}) (*ListLeagueMatches, error) {
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

	res.Decode(&listLeagueMatches)
	return &listLeagueMatches, nil
}
