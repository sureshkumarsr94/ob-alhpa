package core

import "time"

// Scheme [...]
type Scheme struct {
	ID               string    `gorm:"primaryKey;column:id" json:"-"`
	Type             string    `gorm:"column:type" json:"type"`
	Code             string    `gorm:"column:code" json:"code"`
	Name             string    `gorm:"column:name" json:"name"`
	OutletCategoryID string    `gorm:"column:outlet_category_id" json:"outletCategoryId"`
	CategoryID       string    `gorm:"column:category_id" json:"categoryId"`
	Category         Category  `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryID" json:"comCategoryList"`
	ProductID        string    `gorm:"column:product_id" json:"productId"`
	Product          Product   `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	AreaID           string    `gorm:"column:area_id" json:"areaId"`
	CoreArea         Area      `gorm:"joinForeignKey:area_id;foreignKey:id;references:AreaID" json:"coreAreaList"`
	CityID           string    `gorm:"column:city_id" json:"cityId"`
	CoreCity         City      `gorm:"joinForeignKey:city_id;foreignKey:id;references:CityID" json:"coreCityList"`
	StateID          string    `gorm:"column:state_id" json:"stateId"`
	CoreState        State     `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"coreStateList"`
	CountryID        string    `gorm:"column:country_id" json:"countryId"`
	CoreCountry      Country   `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"coreCountryList"`
	StartDate        time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate          time.Time `gorm:"column:end_date" json:"endDate"`
	ItemStart        int       `gorm:"column:item_start" json:"itemStart"`
	ItemEnd          int       `gorm:"column:item_end" json:"itemEnd"`
	ItemDesc         float64   `gorm:"column:item_desc" json:"itemDesc"`
	ItemDiscMax      float64   `gorm:"column:item_disc_max" json:"itemDiscMax"`
	FreeProductID    string    `gorm:"column:free_product_id" json:"freeProductId"`
	CreatedBy        string    `gorm:"column:created_by" json:"createdBy"`
	CreatedByUser    User      `gorm:"joinForeignKey:created_by;foreignKey:id;references:CreatedBy" json:"coreUserList"`
	DistributorID    string    `gorm:"column:distributor_id" json:"distributorId"`
	Status           string    `gorm:"column:status" json:"status"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Scheme) TableName() string {
	return "core_scheme"
}

// SchemeColumns get sql column name.获取数据库列名
var SchemeColumns = struct {
	ID               string
	Type             string
	Code             string
	Name             string
	OutletCategoryID string
	CategoryID       string
	ProductID        string
	AreaID           string
	CityID           string
	StateID          string
	CountryID        string
	StartDate        string
	EndDate          string
	ItemStart        string
	ItemEnd          string
	ItemDesc         string
	ItemDiscMax      string
	FreeProductID    string
	CreatedBy        string
	DistributorID    string
	Status           string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	Type:             "type",
	Code:             "code",
	Name:             "name",
	OutletCategoryID: "outlet_category_id",
	CategoryID:       "category_id",
	ProductID:        "product_id",
	AreaID:           "area_id",
	CityID:           "city_id",
	StateID:          "state_id",
	CountryID:        "country_id",
	StartDate:        "start_date",
	EndDate:          "end_date",
	ItemStart:        "item_start",
	ItemEnd:          "item_end",
	ItemDesc:         "item_desc",
	ItemDiscMax:      "item_disc_max",
	FreeProductID:    "free_product_id",
	CreatedBy:        "created_by",
	DistributorID:    "distributor_id",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}
