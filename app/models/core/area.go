package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Area [...]
type Area struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Pincode     int       `gorm:"column:pincode" json:"pincode"`
	Description string    `gorm:"column:description" json:"description"`
	CityID      string    `gorm:"column:city_id" json:"city_id"`
	City        City      `gorm:"joinForeignKey:city_id;foreignKey:id;references:CityID" json:"city"`
	StateID     string    `gorm:"column:state_id" json:"state_id"`
	State       State     `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"state"`
	CountryID   string    `gorm:"column:country_id" json:"country_id"`
	Country     Country   `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"country"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Area) TableName() string {
	return "core_area"
}

// AreaColumns get sql column name.获取数据库列名
var AreaColumns = struct {
	ID          string
	Name        string
	Pincode     string
	Description string
	CityID      string
	StateID     string
	CountryID   string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Pincode:     "pincode",
	Description: "description",
	CityID:      "city_id",
	StateID:     "state_id",
	CountryID:   "country_id",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

func (m *Area) FindByPrimaryKey(id string) (result Area, err error) {
	err = database.MysqlDB.Model(m).Preload("City").Preload("State").Preload("Country").Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Area) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Area) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Area) FindAll(whereCondition []database.WhereCondition) (results []Area, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Preload("City").Preload("State").Preload("Country").Find(&results).Error
	return
}

func (m *Area) Save() (result Area, err error) {
	err = database.MysqlDB.Save(&m).Preload("City").Preload("State").Preload("Country").Last(&result).Error
	return
}

func (m *Area) FindAllWithGroup(whereCondition []database.WhereCondition) (results []Area, err error) {
	db := database.MysqlDB.Model(m).Preload("City").Preload("State").Preload("Country")
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	db = db.Select("DISTINCT pincode").Limit(10)
	err = db.Find(&results).Error
	return
}
