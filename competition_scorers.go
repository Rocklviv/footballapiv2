package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// CompetitionScorers returns all goal poachers in current season and required league.
// To be able to get goal poachers from different season use filters.
func (c *Client) CompetitionScorers(competitionID int16, values interface{}) (*ListLeagueScorers, error) {
	var scorers ListLeagueScorers
	var res *json.Decoder
	var err error

	path := fmt.Sprintf("competitions/%d/scorers", competitionID)

	if values == nil {
		res, err = c.doRequest("GET", path, nil)
	} else {
		res, err = c.doRequest("GET", path, structToMap(values))
	}

	if err != nil {
		return nil, err
	}

	err = res.Decode(&scorers)
	if err != nil {
		return nil, err
	}
	return &scorers, nil
}
