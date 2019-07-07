package footballapiv2

import "fmt"

// GetCompetition returns one particullar competition.
func (c *Client) GetCompetition(id int) (*Competition, error) {
	var competition Competition
	path := fmt.Sprintf("competitions/%d", id)
	res, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res.Decode(&competition)
	return &competition, nil
}
