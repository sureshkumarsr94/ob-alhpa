package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// OutletCategory [...]
type OutletCategory struct {
	ID            string      `gorm:"primaryKey;column:id" json:"-"`
	Name          string      `gorm:"column:name" json:"name"`
	Description   string      `gorm:"column:description" json:"description"`
	DistributorID null.String `gorm:"column:distributor_id" json:"distributorId"`
	Distributor   Distributor `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	Status        string      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time   `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *OutletCategory) TableName() string {
	return "ent_outlet_category"
}

// OutletCategoryColumns get sql column name.获取数据库列名
var OutletCategoryColumns = struct {
	ID            string
	Name          string
	Description   string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Name:          "name",
	Description:   "description",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *OutletCategory) FindByPrimaryKey(ID string) (result OutletCategory, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", ID).Find(&result).Error
	return
}

func (m *OutletCategory) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *OutletCategory) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *OutletCategory) FindAll(whereCondition []database.WhereCondition) (results []OutletCategory, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *OutletCategory) Save() (result OutletCategory, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *OutletCategory) FindOneByCondition(condition []database.WhereCondition) (result OutletCategory, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &condition, nil, nil)
	err = db.First(&result).Error
	return
}
