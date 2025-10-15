package location_dto

type CountryRequest struct {
	CountryId   string `json:"id"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required"`
}

type StateRequest struct {
	StateId     string `json:"state_id"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CountryId   string `json:"country_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type CityRequest struct {
	CityId      string `json:"city_id"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CountryId   string `json:"country_id" validate:"required"`
	StateId     string `json:"state_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type AreaRequest struct {
	AreaId      string `json:"area_id"`
	Pincode     string `json:"pincode" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CountryId   string `json:"country_id" validate:"required"`
	StateId     string `json:"state_id" validate:"required"`
	CityId      string `json:"city_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
