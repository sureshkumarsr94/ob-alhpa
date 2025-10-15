package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// OutletAddress [...]
type OutletAddress struct {
	ID          string            `gorm:"primaryKey;column:id" json:"id"`
	OutletID    string            `gorm:"column:outlet_id" json:"outletId"`
	Outlet      Outlet            `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	AddressType string            `gorm:"column:address_type" json:"addressType"`
	Address     string            `gorm:"column:address" json:"address"`
	AreaID      string            `gorm:"column:area_id" json:"areaId"`
	CoreArea    coreModel.Area    `gorm:"joinForeignKey:area_id;foreignKey:id;references:AreaID" json:"coreAreaList"`
	CityID      string            `gorm:"column:city_id" json:"cityId"`
	CoreCity    coreModel.City    `gorm:"joinForeignKey:city_id;foreignKey:id;references:CityID" json:"coreCityList"`
	StateID     string            `gorm:"column:state_id" json:"stateId"`
	CoreState   coreModel.State   `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"coreStateList"`
	CountryID   string            `gorm:"column:country_id" json:"countryId"`
	CoreCountry coreModel.Country `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"coreCountryList"`
	Landmark    string            `gorm:"column:landmark" json:"landmark"`
	Pincode     int               `gorm:"column:pincode" json:"pincode"`
	Status      string            `gorm:"column:status" json:"status"`
	CreatedAt   time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m OutletAddress) TableName() string {
	return "ent_outlet_address"
}

// OutletAddressColumns get sql column name.获取数据库列名
var OutletAddressColumns = struct {
	ID          string
	OutletID    string
	AddressType string
	Address     string
	AreaID      string
	CityID      string
	StateID     string
	CountryID   string
	Landmark    string
	Pincode     string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	OutletID:    "outlet_id",
	AddressType: "address_type",
	Address:     "address",
	AreaID:      "area_id",
	CityID:      "city_id",
	StateID:     "state_id",
	CountryID:   "country_id",
	Landmark:    "landmark",
	Pincode:     "pincode",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

func (m OutletAddress) FindByPrimaryKey(ID string) (result OutletAddress, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", ID).Find(&result).Error
	return
}

func (m OutletAddress) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m OutletAddress) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m OutletAddress) FindAll(whereCondition []database.WhereCondition) (results []OutletAddress, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m OutletAddress) Save() (result OutletAddress, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m OutletAddress) FindOneByCondition(condition []database.WhereCondition) (result OutletAddress, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &condition, nil, nil)
	err = db.First(&result).Error
	return
}
