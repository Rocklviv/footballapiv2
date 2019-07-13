package footballapiv2

import "fmt"

// GetCompetition returns one particullar competition.
func (c *Client) GetCompetition(id int16) (*Competition, error) {
	var competition Competition
	path := fmt.Sprintf("competitions/%d", id)
	res, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	err = res.Decode(&competition)
	if err != nil {
		return nil, err
	}
	return &competition, nil
}
