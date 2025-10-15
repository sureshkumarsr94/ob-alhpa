package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Category [...]
type Category struct {
	ID               string      `gorm:"primaryKey;column:id" json:"-"`
	Code             string      `gorm:"column:code" json:"code"`
	Name             string      `gorm:"column:name" json:"name"`
	Description      string      `gorm:"column:description" json:"description"`
	ParentCategoryID null.String `gorm:"column:parent_category_id" json:"parentCategoryId"`
	ParentCategory   *Category   `gorm:"joinForeignKey:parent_category_id;foreignKey:id;references:ParentCategoryID" json:"parentCategory"`
	Sequence         int         `gorm:"column:sequence" json:"sequence"`
	DistributorID    null.String `gorm:"column:distributor_id" json:"distributorId"`
	Status           string      `gorm:"column:status" json:"status"`
	CreatedAt        time.Time   `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        time.Time   `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Category) TableName() string {
	return "core_category"
}

// CategoryColumns get sql column name.获取数据库列名
var CategoryColumns = struct {
	ID               string
	Code             string
	Name             string
	Description      string
	ParentCategoryID string
	Sequence         string
	DistributorID    string
	Status           string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	Code:             "code",
	Name:             "name",
	Description:      "description",
	ParentCategoryID: "parent_category_id",
	Sequence:         "sequence",
	DistributorID:    "distributor_id",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

func (m *Category) FindByPrimaryKey(id string) (result Category, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Category) FindAll(whereCondition []database.WhereCondition) (results []Category, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Category) Save() (result Category, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
