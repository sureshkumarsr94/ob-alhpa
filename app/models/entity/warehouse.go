package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Warehouse [...]
type Warehouse struct {
	ID            string      `gorm:"primaryKey;column:id" json:"id"`
	Code          string      `gorm:"column:code" json:"code"`
	Name          string      `gorm:"column:name" json:"name"`
	Description   string      `gorm:"column:description" json:"description"`
	DistributorID null.String `gorm:"column:distributor_id" json:"distributorId"`
	Status        string      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Warehouse) TableName() string {
	return "ent_warehouse"
}

// WarehouseColumns get sql column name.获取数据库列名
var WarehouseColumns = struct {
	ID            string
	Code          string
	Name          string
	Description   string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Code:          "code",
	Name:          "name",
	Description:   "description",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *Warehouse) FindByPrimaryKey(id string) (result Warehouse, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Warehouse) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Warehouse) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Warehouse) FindAll(whereCondition []database.WhereCondition) (results []Warehouse, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Warehouse) Save() (result Warehouse, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
