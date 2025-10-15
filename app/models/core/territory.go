package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Territory [...]
type Territory struct {
	ID                string        `gorm:"primaryKey;column:id" json:"-"`
	Code              string        `gorm:"column:code" json:"code"`
	Name              string        `gorm:"column:name" json:"name"`
	Description       string        `gorm:"column:description" json:"description"`
	TerritoryTypeID   string        `gorm:"column:territory_type_id" json:"territoryTypeId"`
	TerritoryType     TerritoryType `gorm:"joinForeignKey:territory_type_id;foreignKey:id;references:TerritoryTypeID" json:"coreTerritoryTypeList"`
	ParentID          null.String   `gorm:"column:parent_id" json:"parentId"`
	ParentTerritory   *Territory    `gorm:"joinForeignKey:parent_id;foreignKey:id;references:ParentID" json:"coreTerritoryList"`
	LocationType      null.String   `gorm:"column:location_type" json:"locationType"`
	LocationTypeValue null.String   `gorm:"column:location_type_value" json:"locationTypeValue"`
	DistributorID     null.String   `gorm:"column:distributor_id" json:"distributorId"`
	Status            string        `gorm:"column:status" json:"status"`
	CreatedAt         time.Time     `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time     `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Territory) TableName() string {
	return "core_territory"
}

// TerritoryColumns get sql column name.获取数据库列名
var TerritoryColumns = struct {
	ID                string
	Code              string
	Name              string
	Description       string
	TerritoryTypeID   string
	ParentID          string
	LocationType      string
	LocationTypeValue string
	DistributorID     string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	Code:              "code",
	Name:              "name",
	Description:       "description",
	TerritoryTypeID:   "territory_type_id",
	ParentID:          "parent_id",
	LocationType:      "location_type",
	LocationTypeValue: "location_type_value",
	DistributorID:     "distributor_id",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

func (m *Territory) FindByPrimaryKey(id string) (result Territory, err error) {
	err = database.MysqlDB.Model(m).Preload("TerritoryType").Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Territory) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Territory) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Territory) FindAll(whereCondition []database.WhereCondition) (results []Territory, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Preload("TerritoryType").Find(&results).Error
	return
}

func (m *Territory) Save() (result Territory, err error) {
	err = database.MysqlDB.Save(&m).Preload("TerritoryType").Last(&result).Error
	return
}
