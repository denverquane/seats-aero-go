package seats_aero

import (
	"fmt"
	"time"
)

func TripSearchURL(id string) (string, error) {
	return fmt.Sprintf("%s/%s", TRIP_URL, id), nil
}

type Trip struct {
	Data                   []AvailabilityData `json:"data"`
	OriginCoordinates      Coordinates        `json:"origin_coordinates"`
	DestinationCoordinates Coordinates        `json:"destination_coordinates"`
	BookingLinks           []BookingLink      `json:"booking_links"`
	RevalidationID         string             `json:"revalidation_id"`
}

type Coordinates struct {
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
}

type BookingLink struct {
	Label   string `json:"label"`
	Link    string `json:"link"`
	Primary bool   `json:"primary"`
}

type AvailabilityData struct {
	ID                   string                `json:"ID"`
	RouteID              string                `json:"RouteID"`
	AvailabilityID       string                `json:"AvailabilityID"`
	AvailabilitySegments []AvailabilitySegment `json:"AvailabilitySegments"`
	TotalDuration        int                   `json:"TotalDuration"`
	Stops                int                   `json:"Stops"`
	Carriers             string                `json:"Carriers"`
	RemainingSeats       int                   `json:"RemainingSeats"`
	MileageCost          int                   `json:"MileageCost"`
	TotalTaxes           int                   `json:"TotalTaxes"`
	TaxesCurrency        string                `json:"TaxesCurrency"`
	TaxesCurrencySymbol  string                `json:"TaxesCurrencySymbol"`
	AllianceCost         int                   `json:"AllianceCost"`
	TotalSegmentDistance int                   `json:"TotalSegmentDistance"`
	FlightNumbers        string                `json:"FlightNumbers"`
	DepartsAt            time.Time             `json:"DepartsAt"`
	Cabin                string                `json:"Cabin"`
	ArrivesAt            time.Time             `json:"ArrivesAt"`
	CreatedAt            time.Time             `json:"CreatedAt"`
	UpdatedAt            time.Time             `json:"UpdatedAt"`
	Source               string                `json:"Source"`
	Filtered             bool                  `json:"Filtered"`
}

type AvailabilitySegment struct {
	ID                 string    `json:"ID"`
	RouteID            string    `json:"RouteID"`
	AvailabilityID     string    `json:"AvailabilityID"`
	AvailabilityTripID string    `json:"AvailabilityTripID"`
	FlightNumber       string    `json:"FlightNumber"`
	Distance           int       `json:"Distance"`
	FareClass          string    `json:"FareClass"`
	AircraftName       string    `json:"AircraftName"`
	AircraftCode       string    `json:"AircraftCode"`
	OriginAirport      string    `json:"OriginAirport"`
	DestinationAirport string    `json:"DestinationAirport"`
	DepartsAt          time.Time `json:"DepartsAt"`
	ArrivesAt          time.Time `json:"ArrivesAt"`
	CreatedAt          time.Time `json:"CreatedAt"`
	UpdatedAt          time.Time `json:"UpdatedAt"`
	Source             string    `json:"Source"`
	Cabin              string    `json:"Cabin"`
	Order              int       `json:"Order"`
}
