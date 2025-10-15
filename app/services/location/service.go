package location_service

import (
	"infopack.co.in/offybox/app/dto"
	locationDto "infopack.co.in/offybox/app/dto/location"
	coreModel "infopack.co.in/offybox/app/models/core"
)

// LocationService defines the interface for location-related operations
type LocationService interface {
	// SaveCountry save the country
	SaveCountry(params locationDto.CountryRequest) (country coreModel.Country, handle dto.HandleError)

	// SaveState save the state
	SaveState(params locationDto.StateRequest) (state coreModel.State, handle dto.HandleError)

	// SaveCity save the city
	SaveCity(params locationDto.CityRequest) (city coreModel.City, handle dto.HandleError)

	// SaveArea save the area
	SaveArea(params locationDto.AreaRequest) (area coreModel.Area, handle dto.HandleError)
}

// locationService is an implementation of LocationService
type locationService struct{}

// NewLocationService returns a new instance of LocationService
func NewLocationService() LocationService {
	return &locationService{}
}
