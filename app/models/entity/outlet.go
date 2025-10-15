package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// Outlet [...]
type Outlet struct {
	ID                string         `gorm:"primaryKey;column:id" json:"id"`
	Code              string         `gorm:"column:code" json:"code"`
	Name              string         `gorm:"column:name" json:"name"`
	Type              string         `gorm:"column:type" json:"type"`
	Email             string         `gorm:"column:email" json:"email"`
	Mobile            string         `gorm:"column:mobile" json:"mobile"`
	CreditLimit       float64        `gorm:"column:credit_limit" json:"credit_limit"`
	Outstanding       float64        `gorm:"column:outstanding" json:"outstanding"`
	IncorporationDate time.Time      `gorm:"column:incorporation_date" json:"incorporation_date"`
	CreatedBy         string         `gorm:"column:created_by" json:"created_by"`
	CoreUser          coreModel.User `gorm:"joinForeignKey:created_by;foreignKey:id;references:CreatedBy" json:"created_by_user"`
	OutletCategoryID  string         `gorm:"column:outlet_category_id" json:"outlet_category_id"`
	OutletCategory    OutletCategory `gorm:"joinForeignKey:outlet_category_id;foreignKey:id;references:OutletCategoryID" json:"outlet_category"`
	DistributorID     null.String    `gorm:"column:distributor_id" json:"distributor_id"`
	Distributor       Distributor    `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"distributor"`
	Status            string         `gorm:"column:status" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Outlet) TableName() string {
	return "ent_outlet"
}

// OutletColumns get sql column name.获取数据库列名
var OutletColumns = struct {
	ID                string
	Code              string
	Name              string
	Type              string
	Email             string
	Mobile            string
	CreditLimit       string
	Outstanding       string
	IncorporationDate string
	CreatedBy         string
	OutletID          string
	DistributorID     string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	Code:              "code",
	Name:              "name",
	Type:              "type",
	Email:             "email",
	Mobile:            "mobile",
	CreditLimit:       "credit_limit",
	Outstanding:       "outstanding",
	IncorporationDate: "incorporation_date",
	CreatedBy:         "created_by",
	OutletID:          "outlet_category_id",
	DistributorID:     "distributor_id",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

func (m *Outlet) FindByPrimaryKey(ID string) (result Outlet, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", ID).Find(&result).Error
	return
}

func (m *Outlet) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Outlet) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Outlet) FindAll(whereCondition []database.WhereCondition) (results []Outlet, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Outlet) Save() (result Outlet, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *Outlet) FindOneByCondition(condition []database.WhereCondition) (result Outlet, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &condition, nil, nil)
	err = db.First(&result).Error
	return
}
