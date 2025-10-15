package journey_plan_service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/dto"
	journeyPlanDto "infopack.co.in/offybox/app/dto/journey_plan"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	saleModel "infopack.co.in/offybox/app/models/sale"
	outletService "infopack.co.in/offybox/app/services/outlet"
	userService "infopack.co.in/offybox/app/services/user"
	"time"
)

func (s *journeyPlanService) GetJourneyPlan(journeyPlanID string) (payload journeyPlanDto.JourneyPlanResponse) {
	journeyPlan := saleModel.JourneyPlan{}
	journeyPlan, _ = journeyPlan.FindByPrimaryKey(journeyPlanID)
	payload.JourneyPlanID = journeyPlan.ID

	payload.OutletID = journeyPlan.OutletID
	outletSvc := outletService.NewOutletService()
	userSvc := userService.NewUserService()
	payload.Outlet = outletSvc.GetOutlet(journeyPlan.OutletID, userSvc)

	payload.UserID = journeyPlan.UserID
	user := coreModel.User{}
	payload.User, _ = user.FindByPrimaryKey(journeyPlan.UserID)

	payload.AssignedDate = journeyPlan.AssignedDate.String()
	payload.ClosedDate = journeyPlan.ClosedDate.Time.String()
	payload.Remarks = journeyPlan.Remarks
	payload.Status = journeyPlan.Status

	return
}

func (s *journeyPlanService) SaveJourneyPlan(params journeyPlanDto.JourneyPlanRequest) (journeyPlanID string, handle dto.HandleError) {

	journeyPlan := saleModel.JourneyPlan{}
	if params.JourneyPlanID != "" {
		journeyPlan, _ = journeyPlan.FindByPrimaryKey(params.JourneyPlanID)
		if journeyPlan.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("jounrney plan is not found")
			return
		}
	} else {
		journeyPlan.ID = uuid.New().String()
		journeyPlan.CreatedAt = time.Now()
		journeyPlan.UpdatedAt = time.Now()
	}

	journeyPlan.OutletID = params.OutletID
	outlet := entityModel.Outlet{}
	outlet, _ = outlet.FindByPrimaryKey(params.OutletID)
	if len(outlet.ID) == 0 {
		handle.Status = -2
		handle.Errors = fmt.Errorf("outlet is not found")
		return
	}
	journeyPlan.UserID = params.UserID
	user := coreModel.User{}
	user, _ = user.FindByPrimaryKey(journeyPlan.UserID)
	if len(user.ID) == 0 {
		handle.Status = -3
		handle.Errors = fmt.Errorf("user is not found")
		return
	}

	journeyPlan.AssignedDate = utility.ParseDate(params.AssignedDate)
	journeyPlan.ClosedDate = null.TimeFrom(utility.ParseDate(params.ClosedDate))
	journeyPlan.Remarks = params.Remarks
	journeyPlan.Status = params.Status
	journeyPlan, err := journeyPlan.Save()
	if err != nil {
		handle = dto.HandleError{
			Status: -2,
			Errors: err,
		}
		return
	}
	journeyPlanID = journeyPlan.ID
	return
}
