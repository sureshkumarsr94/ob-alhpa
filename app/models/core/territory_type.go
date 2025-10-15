package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// TerritoryType [...]
type TerritoryType struct {
	ID                  string         `gorm:"primaryKey;column:id" json:"-"`
	Code                string         `gorm:"column:code" json:"code"`
	Name                string         `gorm:"column:name" json:"name"`
	Description         string         `gorm:"column:description" json:"description"`
	Level               int            `gorm:"column:level" json:"level"`
	ParentID            null.String    `gorm:"column:parent_id" json:"parentId"`
	ParentTerritoryType *TerritoryType `gorm:"joinForeignKey:parent_id;foreignKey:id;references:ParentID" json:"coreTerritoryTypeList"`
	DistributorID       null.String    `gorm:"column:distributor_id" json:"distributorId"`
	Status              string         `gorm:"column:status" json:"status"`
	CreatedAt           time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.
func (m *TerritoryType) TableName() string {
	return "core_territory_type"
}

// TerritoryTypeColumns get sql column name.获取数据库列名
var TerritoryTypeColumns = struct {
	ID            string
	Code          string
	Name          string
	Description   string
	Level         string
	ParentID      string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Code:          "code",
	Name:          "name",
	Description:   "description",
	Level:         "level",
	ParentID:      "parent_id",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *TerritoryType) FindByPrimaryKey(id string) (result TerritoryType, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *TerritoryType) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *TerritoryType) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *TerritoryType) FindAll(whereCondition []database.WhereCondition) (results []TerritoryType, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *TerritoryType) Save() (result TerritoryType, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
