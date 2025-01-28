# seats-aero-go
Unofficial Go library for accessing the [seats.aero API](https://developers.seats.aero/reference/getting-started-p)

Requires a Pro subscription (and associated API Key) from seats.aero, which you can find under the API header in your [seats.aero account settings](https://seats.aero/settings)

# API Implementation Status
- [X] Cached Search
- [ ] Bulk Availability
- [X] Get Trips
- [ ] Get Routes
- [ ] Live Search (Commercial Agreement needed; probably won't implement)

# Installation

`go get github.com/denverquane/seats-aero-go`

# Usage
```
import (
    ...
    "github.com/denverquane/ana-rtw-monitor/seats_aero"
)

func main() {
    seatsApiKey := os.Getenv("SEATS_AERO_API_KEY")

    seatsAeroClient := seats_aero.New(seatsApiKey)
    
    cachedSearchResult, err := seatsAeroClient.CachedSearch("IAD,DFW", "IST", "business", "2025-01-01", "2025-01-02")
    if err != nil {
        log.Fatal(err)
    }
    
    for _, data := range cachedSearchResult.Data {
    
        // CAUTION: filter the data here based on criteria you care about before doing a trip search, 
        // like ensuring there's business class seats available, United shows availability, etc 
        // not doing so will consume a ton of your limited daily API requests!
        
        trip, err := seatsAeroClient.TripSearch(data.ID)
        if err != nil {
            log.Println(err)
        }
        
        ...
    }
}
```
