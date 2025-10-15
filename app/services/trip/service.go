package trip_service

import (
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/dto"
	tripDto "infopack.co.in/offybox/app/dto/trip"
	coreModel "infopack.co.in/offybox/app/models/core"
)

// TripService defines the interface for trip-related operations
type TripService interface {

	// GetTrip retrieves the territory type object
	GetTrip(tripID string) (payload tripDto.TripResponse)

	// SaveTrip save the trip
	SaveTrip(tx *gorm.DB, params tripDto.TripRequest, userObject coreModel.User) (tripID string, handle dto.HandleError)
}

// tripService is an implementation of TripService
type tripService struct{}

// NewTripService returns a new instance of TripService
func NewTripService() TripService {
	return &tripService{}
}
