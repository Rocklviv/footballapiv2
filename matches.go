package footballapiv2

import (
	"encoding/json"
	"fmt"
)

// ListOfMatches returns all matches for selected competition(s).
func (c *Client) ListOfMatches(values interface{}) (*ListMatches, error) {
	var matches ListMatches
	var res *json.Decoder
	var err error

	path := fmt.Sprintf("matches")

	if values == nil {
		res, err = c.doRequest("GET", path, nil)
	} else {
		if validateFilter(values, "dateFrom") {
			if validateFilter(values, "dateFrom") && validateFilter(values, "dateTo") {
				res, err = c.doRequest("GET", path, structToMap(values))
			} else {
				return nil, fmt.Errorf("Filter(s) dateFrom cannot be used without dateTo same as dateTo cannot be used without dateFrom")
			}
		} else {
			res, err = c.doRequest("GET", path, structToMap(values))
		}
	}

	if err != nil {
		return nil, err
	}

	res.Decode(&matches)
	return &matches, nil
}

// GetMatch returns one required match.
func (c *Client) GetMatch(matchID uint32) (*Head2HeadMatch, error) {
	var match Head2HeadMatch
	if matchID != 0 {
		path := fmt.Sprintf("matches/%d", matchID)

		res, err := c.doRequest("GET", path, nil)

		if err != nil {
			return nil, err
		}
		res.Decode(&match)
		return &match, nil
	}
	return nil, fmt.Errorf("matchID should be specified")
}
