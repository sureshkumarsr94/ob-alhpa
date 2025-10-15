package journey_plan_service

import (
	"infopack.co.in/offybox/app/dto"
	journeyPlanDto "infopack.co.in/offybox/app/dto/journey_plan"
)

// JourneyPlanService defines the interface for journey plan-related operations
type JourneyPlanService interface {

	// GetJourneyPlan retrieves the  journey plan object
	GetJourneyPlan(journeyPlanID string) (payload journeyPlanDto.JourneyPlanResponse)

	// SaveJourneyPlan save the journey plan
	SaveJourneyPlan(params journeyPlanDto.JourneyPlanRequest) (journeyPlanID string, handle dto.HandleError)
}

// journeyPlanService is an implementation of JourneyPlanService
type journeyPlanService struct{}

// NewJourneyPlanService returns a new instance of JourneyPlanService
func NewJourneyPlanService() JourneyPlanService {
	return &journeyPlanService{}
}
