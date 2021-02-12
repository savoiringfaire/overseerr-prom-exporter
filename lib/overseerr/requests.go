package overseerr

import "encoding/json"

// RequestsResponse contains the response from the request/count endpoint of overseerr API.
type RequestsResponse struct {
	Pending  uint
	Approved uint
}

// RequestCount gets the number of pending and approved requests.
func (c *Overseerr) RequestCount() (*RequestsResponse, error) {
	resp, err := c.makeAPIRequest("request/count")
	if err != nil {
		return nil, err
	}

	response := RequestsResponse{}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
