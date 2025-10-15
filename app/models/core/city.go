package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// City [...]
type City struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	StateID     string    `gorm:"column:state_id" json:"state_id"`
	State       State     `gorm:"joinForeignKey:state_id;foreignKey:id;references:StateID" json:"state"`
	CountryID   string    `gorm:"column:country_id" json:"country_id"`
	Country     Country   `gorm:"joinForeignKey:country_id;foreignKey:id;references:CountryID" json:"country"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *City) TableName() string {
	return "core_city"
}

// CityColumns get sql column name.获取数据库列名
var CityColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	StateID     string
	CountryID   string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Description: "description",
	StateID:     "state_id",
	CountryID:   "country_id",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

func (m *City) FindByPrimaryKey(id string) (result City, err error) {
	err = database.MysqlDB.Model(m).Preload("State").Preload("Country").Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *City) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *City) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *City) FindAll(whereCondition []database.WhereCondition) (results []City, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Preload("State").Preload("Country").Find(&results).Error
	return
}

func (m *City) Save() (result City, err error) {
	err = database.MysqlDB.Save(&m).Preload("State").Preload("Country").Last(&result).Error
	return
}
