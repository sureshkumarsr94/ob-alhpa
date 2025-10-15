package journey_plan_dto

import (
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	coreModel "infopack.co.in/offybox/app/models/core"
)

type JourneyPlanRequest struct {
	JourneyPlanID string `json:"journey_plan_id"`
	OutletID      string `json:"outlet_id" validate:"required"`
	UserID        string `json:"user_id" validate:"required"`
	AssignedDate  string `json:"assigned_date" validate:"required"`
	ClosedDate    string `json:"closed_date"`
	Remarks       string `json:"remarks"`
	Status        string `json:"status"`
}

type JourneyPlanResponse struct {
	JourneyPlanID string                   `json:"journey_plan_id"`
	OutletID      string                   `json:"outlet_id"`
	Outlet        outletDto.OutletResponse `json:"outlet"`
	UserID        string                   `json:"user_id"`
	User          coreModel.User           `json:"user"`
	AssignedDate  string                   `json:"assigned_date"`
	ClosedDate    string                   `json:"closed_date"`
	Remarks       string                   `json:"remarks"`
	Status        string                   `json:"status"`
}
