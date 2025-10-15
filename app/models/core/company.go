package core

import (
	"gorm.io/datatypes"
	"time"
)

// Company [...]
type Company struct {
	ID                string         `gorm:"primaryKey;column:id" json:"-"`
	Name              string         `gorm:"column:name" json:"name"`
	Email             string         `gorm:"column:email" json:"email"`
	Website           string         `gorm:"column:website" json:"website"`
	Mobile            string         `gorm:"column:mobile" json:"mobile"`
	IncorporationDate datatypes.Date `gorm:"column:incorporation_date" json:"incorporationDate"`
	GstNo             string         `gorm:"column:gst_no" json:"gstNo"`
	Address           string         `gorm:"column:address" json:"address"`
	AreaID            string         `gorm:"column:area_id" json:"areaId"`
	Area              Area           `gorm:"joinForeignKey:area_id;foreignKey:id;references:AreaID" json:"coreAreaList"`
	CityID            string         `gorm:"column:city_id" json:"cityId"`
	City              City           `gorm:"joinForeignKey:city_id;foreignKey:id;references:CityID" json:"coreCityList"`
	StateID           string         `gorm:"column:state_id" json:"stateId"`
	State             State          `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"coreStateList"`
	CountryID         string         `gorm:"column:country_id" json:"countryId"`
	Country           Country        `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"coreCountryList"`
	Status            string         `gorm:"column:status" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Company) TableName() string {
	return "core_company"
}

// CompanyColumns get sql column name.获取数据库列名
var CompanyColumns = struct {
	ID                string
	Name              string
	Email             string
	Website           string
	Mobile            string
	IncorporationDate string
	GstNo             string
	Address           string
	AreaID            string
	CityID            string
	StateID           string
	CountryID         string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	Name:              "name",
	Email:             "email",
	Website:           "website",
	Mobile:            "mobile",
	IncorporationDate: "incorporation_date",
	GstNo:             "gst_no",
	Address:           "address",
	AreaID:            "area_id",
	CityID:            "city_id",
	StateID:           "state_id",
	CountryID:         "country_id",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}
