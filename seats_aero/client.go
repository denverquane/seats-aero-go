package seats_aero

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	API_URL    = "https://seats.aero/partnerapi"
	SEARCH_URL = API_URL + "/search"
	TRIP_URL   = API_URL + "/trips"
)

const AirportRegex = `^([A-Z]{3})(,[A-Z]{3})*$`

type Cabin string

const (
	Economy  Cabin = "economy"
	Business Cabin = "business"
	First    Cabin = "first"
)

type SeatsAeroClient struct {
	apiKey string
}

func New(apiKey string) *SeatsAeroClient {
	return &SeatsAeroClient{apiKey: apiKey}
}

func (c *SeatsAeroClient) CachedSearch(origin, dest string, cabin Cabin, startDate, endDate string) (CachedSearchResponse, error) {
	var result CachedSearchResponse

	url, err := CachedSearchURL(origin, dest, cabin, startDate, endDate)
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Partner-Authorization", c.apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, errors.New("GET " + url + " failed with response: " + res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	body := string(bodyBytes)

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *SeatsAeroClient) TripSearch(id string) (Trip, error) {
	var result Trip

	url, err := TripSearchURL(id)
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Partner-Authorization", c.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, errors.New("GET " + url + " failed with response: " + res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	body := string(bodyBytes)

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
