package location_service

import (
	"fmt"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/dto"
	locationDto "infopack.co.in/offybox/app/dto/location"
	coreModel "infopack.co.in/offybox/app/models/core"
)

func (s *locationService) SaveCountry(params locationDto.CountryRequest) (country coreModel.Country, handle dto.HandleError) {
	if params.CountryId != "" {
		country, _ = country.FindByPrimaryKey(params.CountryId)
		if country.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid country")
			return
		}
	}
	country.Code = params.Code
	country.Name = params.Name
	country.Description = params.Description
	country.Status = params.Status
	country, err := country.Save()
	if err != nil {
		handle.Status = -2
		handle.Errors = err
		return
	}
	return
}

func (s *locationService) SaveState(params locationDto.StateRequest) (state coreModel.State, handle dto.HandleError) {
	if params.CountryId != "" {
		country := coreModel.Country{}
		country, _ = country.FindByPrimaryKey(params.CountryId)
		if country.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid country")
			return
		}
	}
	if params.StateId != "" {
		state, _ = state.FindByPrimaryKey(params.StateId)
		if state.ID == "" {
			handle.Status = -2
			handle.Errors = fmt.Errorf("invalid state")
			return
		}
	}
	state.Code = params.Code
	state.Name = params.Name
	state.Description = params.Description
	state.Status = params.Status
	state.CountryID = params.CountryId
	state, err := state.Save()
	if err != nil {
		handle.Status = -3
		handle.Errors = err
		return
	}
	return
}

func (s *locationService) SaveCity(params locationDto.CityRequest) (city coreModel.City, handle dto.HandleError) {

	if params.CountryId != "" {
		country := coreModel.Country{}
		country, _ = country.FindByPrimaryKey(params.CountryId)
		if country.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid country")
			return
		}
	}
	if params.StateId != "" {
		state := coreModel.State{}
		state, _ = state.FindByPrimaryKey(params.StateId)
		if state.ID == "" {
			handle.Status = -2
			handle.Errors = fmt.Errorf("invalid state")
			return
		}
	}
	if params.CityId != "" {
		city, _ = city.FindByPrimaryKey(params.CityId)
		if city.ID == "" {
			handle.Status = -3
			handle.Errors = fmt.Errorf("invalid city")
			return
		}
	}
	city.Code = params.Code
	city.Name = params.Name
	city.Description = params.Description
	city.Status = params.Status
	city.CountryID = params.CountryId
	city.StateID = params.StateId
	city, err := city.Save()
	if err != nil {
		handle.Status = -4
		handle.Errors = err
		return
	}
	return
}

func (s *locationService) SaveArea(params locationDto.AreaRequest) (area coreModel.Area, handle dto.HandleError) {

	if params.CountryId != "" {
		country := coreModel.Country{}
		country, _ = country.FindByPrimaryKey(params.CountryId)
		if country.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid country")
			return
		}
	}
	if params.StateId != "" {
		state := coreModel.State{}
		state, _ = state.FindByPrimaryKey(params.StateId)
		if state.ID == "" {
			handle.Status = -2
			handle.Errors = fmt.Errorf("invalid state")
			return
		}
	}
	if params.CityId != "" {
		city := coreModel.City{}
		city, _ = city.FindByPrimaryKey(params.CityId)
		if city.ID == "" {
			handle.Status = -3
			handle.Errors = fmt.Errorf("invalid city")
			return
		}
	}
	if params.AreaId != "" {
		area, _ = area.FindByPrimaryKey(params.AreaId)
		if area.ID == "" {
			handle.Status = -4
			handle.Errors = fmt.Errorf("invalid area")
			return
		}
	}
	area.Pincode = utility.StringToInt(params.Pincode)
	area.Name = params.Name
	area.Description = params.Description
	area.Status = params.Status
	area.CountryID = params.CountryId
	area.StateID = params.StateId
	area.CityID = params.CityId
	area, err := area.Save()
	if err != nil {
		handle.Status = -5
		handle.Errors = err
		return
	}
	return
}
