package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreLocation -
func (c *Client) ExploreLocation(pageURL string) (Location, error) {
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}
