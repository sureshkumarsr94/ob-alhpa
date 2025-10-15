package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Product [...]
type Product struct {
	ID            string      `gorm:"primaryKey;column:id" json:"id"`
	Code          string      `gorm:"column:code" json:"code"`
	Name          string      `gorm:"column:name" json:"name"`
	Mrp           float64     `gorm:"column:mrp" json:"mrp"`
	Uom           string      `gorm:"column:uom" json:"uom"`
	Variant       string      `gorm:"column:variant" json:"variant"`
	Description   string      `gorm:"column:description" json:"description"`
	Barcode       string      `gorm:"column:barcode" json:"barcode"`
	DistributorID null.String `gorm:"column:distributor_id" json:"distributor_id"`
	Status        string      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "core_product"
}

// ProductColumns get sql column name.获取数据库列名
var ProductColumns = struct {
	ID            string
	Code          string
	Name          string
	Mrp           string
	Uom           string
	Variant       string
	Description   string
	Barcode       string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Code:          "code",
	Name:          "name",
	Mrp:           "mrp",
	Uom:           "uom",
	Variant:       "variant",
	Description:   "description",
	Barcode:       "barcode",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *Product) FindByPrimaryKey(id string) (result Product, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Product) FindAll(whereCondition []database.WhereCondition) (results []Product, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Product) FindOneWithCondition(whereCondition []database.WhereCondition) (results Product, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Product) Save() (result Product, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
