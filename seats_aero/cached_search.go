package seats_aero

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func CachedSearchURL(origin, dest string, cabin Cabin, startDate, endDate string) (string, error) {
	_, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return "", err
	}
	_, err = time.Parse("2006-01-02", endDate)
	if err != nil {
		return "", err
	}

	if cabin != Economy && cabin != Business && cabin != First {
		return "", errors.New(fmt.Sprintf("cabin should be '%s', '%s', or '%s'", Economy, Business, First))
	}
	origin = strings.ToUpper(origin)
	dest = strings.ToUpper(dest)

	_, err = regexp.MatchString(AirportRegex, origin)
	if err != nil {
		return "", err
	}
	_, err = regexp.MatchString(AirportRegex, dest)
	if err != nil {
		return "", err
	}
	origin = url.PathEscape(origin)
	dest = url.PathEscape(dest)
	return fmt.Sprintf("%s?origin_airport=%s&destination_airport=%s&cabin=%s&start_date=%s&end_date=%s&take=500",
		SEARCH_URL, origin, dest, cabin, startDate, endDate), nil
}

type CachedSearchResponse struct {
	Data    []CachedSearchData `json:"data"`
	Count   int                `json:"count"`
	HasMore bool               `json:"hasMore"`
	Cursor  int64              `json:"cursor"`
}

type CachedSearchData struct {
	ID                string    `json:"ID"`
	RouteID           string    `json:"RouteID"`
	Route             Route     `json:"Route"`
	Date              string    `json:"Date"`
	ParsedDate        time.Time `json:"ParsedDate"`
	YAvailable        bool      `json:"YAvailable"`
	WAvailable        bool      `json:"WAvailable"`
	JAvailable        bool      `json:"JAvailable"`
	FAvailable        bool      `json:"FAvailable"`
	YMileageCost      string    `json:"YMileageCost"`
	WMileageCost      string    `json:"WMileageCost"`
	JMileageCost      string    `json:"JMileageCost"`
	FMileageCost      string    `json:"FMileageCost"`
	YMileageCostRaw   int       `json:"YMileageCostRaw"`
	WMileageCostRaw   int       `json:"WMileageCostRaw"`
	JMileageCostRaw   int       `json:"JMileageCostRaw"`
	FMileageCostRaw   int       `json:"FMileageCostRaw"`
	TaxesCurrency     string    `json:"TaxesCurrency"`
	YTotalTaxes       int       `json:"YTotalTaxes"`
	WTotalTaxes       int       `json:"WTotalTaxes"`
	JTotalTaxes       int       `json:"JTotalTaxes"`
	FTotalTaxes       int       `json:"FTotalTaxes"`
	YRemainingSeats   int       `json:"YRemainingSeats"`
	WRemainingSeats   int       `json:"WRemainingSeats"`
	JRemainingSeats   int       `json:"JRemainingSeats"`
	FRemainingSeats   int       `json:"FRemainingSeats"`
	YAirlines         string    `json:"YAirlines"`
	WAirlines         string    `json:"WAirlines"`
	JAirlines         string    `json:"JAirlines"`
	FAirlines         string    `json:"FAirlines"`
	Source            string    `json:"Source"`
	CreatedAt         time.Time `json:"CreatedAt"`
	UpdatedAt         time.Time `json:"UpdatedAt"`
	AvailabilityTrips string    `json:"AvailabilityTrips"`
}

type Route struct {
	ID                 string `json:"ID"`
	OriginAirport      string `json:"OriginAirport"`
	OriginRegion       string `json:"OriginRegion"`
	DestinationAirport string `json:"DestinationAirport"`
	DestinationRegion  string `json:"DestinationRegion"`
	NumDaysOut         int    `json:"NumDaysOut"`
	Distance           int    `json:"Distance"`
	Source             string `json:"Source"`
}
