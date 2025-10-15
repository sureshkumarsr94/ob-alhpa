package core

import (
	"gorm.io/datatypes"
	"time"
)

// ProductPrice [...]
type ProductPrice struct {
	ID            string         `gorm:"primaryKey;column:id" json:"-"`
	ProductID     string         `gorm:"column:product_id" json:"productId"`
	Product       Product        `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	AreaID        string         `gorm:"column:area_id" json:"areaId"`
	Area          Area           `gorm:"joinForeignKey:area_id;foreignKey:id;references:AreaID" json:"coreAreaList"`
	CityID        string         `gorm:"column:city_id" json:"cityId"`
	City          City           `gorm:"joinForeignKey:city_id;foreignKey:id;references:CityID" json:"coreCityList"`
	StateID       string         `gorm:"column:state_id" json:"stateId"`
	State         State          `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"coreStateList"`
	CountryID     string         `gorm:"column:country_id" json:"countryId"`
	Country       Country        `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"coreCountryList"`
	StartDate     datatypes.Date `gorm:"column:start_date" json:"startDate"`
	EndDate       datatypes.Date `gorm:"column:end_date" json:"endDate"`
	BatchNo       string         `gorm:"column:batch_no" json:"batchNo"`
	CreatedBy     string         `gorm:"column:created_by" json:"createdBy"`
	CreatedByUser User           `gorm:"joinForeignKey:created_by;foreignKey:id;references:CreatedBy" json:"coreUserList"`
	Status        string         `gorm:"column:status" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ProductPrice) TableName() string {
	return "core_product_price"
}

// ProductPriceColumns get sql column name.获取数据库列名
var ProductPriceColumns = struct {
	ID        string
	ProductID string
	AreaID    string
	CityID    string
	StateID   string
	CountryID string
	StartDate string
	EndDate   string
	BatchNo   string
	CreatedBy string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ProductID: "product_id",
	AreaID:    "area_id",
	CityID:    "city_id",
	StateID:   "state_id",
	CountryID: "country_id",
	StartDate: "start_date",
	EndDate:   "end_date",
	BatchNo:   "batch_no",
	CreatedBy: "created_by",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
