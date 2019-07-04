package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// GetCompetitions get all competitions that are available in api.
func (c *Client) GetCompetitions(values interface{}) []Competition {
	var competitions Competitions
	var res *json.Decoder
	var err error
	if values == nil {
		res, err = c.doRequest("GET", "competitions", nil)
	} else {
		res, err = c.doRequest("GET", "competitions", structToMap(values))
	}
	if err != nil {
		fmt.Println(err)
	}

	err = res.Decode(&competitions)
	if err != nil {
		fmt.Println(err)
	}

	return competitions.Competitions
}
