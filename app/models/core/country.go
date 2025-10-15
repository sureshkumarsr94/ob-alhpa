package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Country [...]
type Country struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Country) TableName() string {
	return "core_country"
}

// CountryColumns get sql column name.获取数据库列名
var CountryColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Description: "description",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

func (m *Country) FindByPrimaryKey(id string) (result Country, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Country) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Country) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Country) FindAll(whereCondition []database.WhereCondition) (results []Country, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Country) Save() (result Country, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
