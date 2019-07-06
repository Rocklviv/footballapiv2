package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// CompetitionStandings returns
func (c *Client) CompetitionStandings(competitionID int, values interface{}) (*ListStandings, error) {
	var standing ListStandings
	var res *json.Decoder
	var err error
	path := fmt.Sprintf("competitions/%d/standings", competitionID)

	if values == nil {
		res, err = c.doRequest("GET", path, nil)
	} else {
		res, err = c.doRequest("GET", path, structToMap(values))
	}

	if err != nil {
		return nil, err
	}

	res.Decode(&standing)
	return &standing, nil
}