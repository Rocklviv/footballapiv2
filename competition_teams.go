package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// CompetitionTeams returns list of teams in competition/league.
func (c *Client) CompetitionTeams(competitionID int, values interface{}) (*TeamsList, error) {
	var teams TeamsList
	var res *json.Decoder
	var err error
	path := fmt.Sprintf("competitions/%d/teams", competitionID)

	if values == nil {
		res, err = c.doRequest("GET", path, nil)
	} else {
		res, err = c.doRequest("GET", path, structToMap(values))
	}

	if err != nil {
		return nil, err
	}

	err = res.Decode(&teams)
	if err != nil {
		return nil, err
	}
	return &teams, nil
}
