package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// GetListOfScorers returns all goal poachers in current season and required league.
// To be able to get goal poachers from different season use filters.
func (c *Client) GetListOfScorers(competitionID uint16, values interface{}) (*ListLeagueScorers, error) {
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

	res.Decode(&scorers)
	return &scorers, nil
}
