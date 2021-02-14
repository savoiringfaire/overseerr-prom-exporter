package overseerr

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RequestsResponse contains the response from the request/count endpoint of overseerr API.
type SearchResponseRaw struct {
	Page         uint          `json:'page'`
	TotalPages   uint          `json:'totalPages'`
	TotalResults uint          `json:'totalResults'`
	Results      []interface{} `json:'results'`
}

type SearchResponse struct {
	Page          uint
	TotalPages    uint
	TotalResults  uint
	TVResults     []SearchResponseTV
	MovieResults  []SearchResponseMovie
	PersonResults []SearchResponsePerson
}

type SearchResponseTV struct {
}

type SearchResponseMovie struct {
	ID               uint `json:'id'`
	MediaType        string
	Popularity       uint
	PosterPath       string
	BackdropPath     string
	VoteCount        uint
	VoteAverage      uint
	GenreIDs         []uint `json:'genreIds'`
	Overview         string
	OriginalLanguage string
	ReleaseDate      string
	Adult            bool
	Video            bool
	MediaInfo        MediaInfo
}

type MediaInfo struct {
	ID        uint `json:'id'`
	TmdbID    uint `json:'tmdbId'`
	TvdbID    uint `json:'tvdbId'`
	Status    uint
	CreatedAt string
	UpdatedAt string
}

type SearchResponsePerson struct {
}

type SearchResponseResults struct {
}

type SearchRequestOptions struct {
	Page     uint
	Language string
}

// RequestCount gets the number of pending and approved requests.
func (c *Overseerr) Search(query string, options SearchRequestOptions) (*SearchResponse, error) {
	if options.Page == 0 {
		options.Page = 1
	}

	resp, err := c.makeAPIRequest(fmt.Sprintf("search?query=%s&page=%d", query, options.Page))
	if err != nil {
		return nil, err
	}

	response := SearchResponseRaw{}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	realResponse := SearchResponse{}

	realResponse.Page = response.Page
	realResponse.TotalPages = response.TotalPages
	realResponse.TotalResults = response.TotalResults

	for _, item := range response.Results {
		resultMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("error")
		}

		switch resultMap["mediaType"] {
		case "movie":
			jsonString, _ := json.Marshal(item)
			fmt.Println(string(jsonString))
			s := SearchResponseMovie{}
			json.Unmarshal(jsonString, &s)
			realResponse.MovieResults = append(realResponse.MovieResults, s)
		}
	}

	return &realResponse, nil
}
