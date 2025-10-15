package outlet_dto

import (
	userDto "infopack.co.in/offybox/app/dto/user"
	coreModel "infopack.co.in/offybox/app/models/core"
)

type OutletCategoryRequest struct {
	OutletCategoryID string `json:"outlet_category_id"`
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	Status           string `json:"status" validate:"required"`
}

type OutletRequest struct {
	OutletID         string  `json:"outlet_id"`
	Code             string  `json:"code" validate:"required"`
	Name             string  `json:"name" validate:"required"`
	Email            string  `json:"email"`
	Mobile           string  `json:"mobile" validate:"required"`
	CreditLimit      float64 `json:"credit_limit" validate:"required"`
	OutletCategoryID string  `json:"outlet_category_id" validate:"required"`
	Status           string  `json:"status" validate:"required"`
}

type OutletAddressRequest struct {
	OutletAddressID string `json:"outlet_address_id"`
	OutletID        string `json:"outlet_id" validate:"required"`
	AddressType     string `json:"address_type" validate:"required"`
	Address         string `json:"address" validate:"required"`
	Landmark        string `json:"landmark"`
	Pincode         string `json:"pincode" validate:"required"`
	AreaId          string `json:"area_id" validate:"required"`
	Status          string `json:"status" validate:"required"`
}

type OutletCategoryResponse struct {
	OutletCategoryID string `json:"outlet_category_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Status           string `json:"status"`
}

type OutletResponse struct {
	OutletID          string                 `json:"outlet_id"`
	Code              string                 `json:"code"`
	Name              string                 `json:"name"`
	Type              string                 `json:"type"`
	Email             string                 `json:"email"`
	Mobile            string                 `json:"mobile"`
	IncorporationDate string                 `json:"incorporation_date"`
	CreditLimit       float64                `json:"credit_limit"`
	Outstanding       float64                `json:"outstanding"`
	CreatedByUserID   string                 `json:"created_by_user_id"`
	CreatedByUser     userDto.UserObject     `json:"created_by_user"`
	OutletCategoryID  string                 `json:"outlet_category_id"`
	OutletCategory    OutletCategoryResponse `json:"outlet_category"`
	Status            string                 `json:"status"`
}

type OutletAddressResponse struct {
	OutletAddressID string            `json:"outlet_address_id"`
	OutletID        string            `json:"outlet_id"`
	AddressType     string            `json:"address_type"`
	Address         string            `json:"address"`
	Landmark        string            `json:"landmark"`
	Pincode         string            `json:"pincode"`
	AreaId          string            `json:"area_id"`
	Area            coreModel.Area    `json:"area" `
	CityId          string            `json:"city_id"`
	City            coreModel.City    `json:"city"`
	StateId         string            `json:"state_id"`
	State           coreModel.State   `json:"state"`
	CountryId       string            `json:"country_id"`
	Country         coreModel.Country `json:"country"`
	Status          string            `json:"status"`
}
