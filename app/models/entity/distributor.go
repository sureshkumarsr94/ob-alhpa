package entity

import (
	"infopack.co.in/offybox/app/database"
	"time"
)

// Distributor [...]
type Distributor struct {
	ID             string    `gorm:"primaryKey;column:id" json:"-"`
	Code           string    `gorm:"column:code" json:"code"`
	Name           string    `gorm:"column:name" json:"name"`
	Email          string    `gorm:"column:email" json:"email"`
	Mobile         string    `gorm:"column:mobile" json:"mobile"`
	PointOfContact string    `gorm:"column:point_of_contact" json:"pointOfContact"`
	Status         string    `gorm:"column:status" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Distributor) TableName() string {
	return "ent_distributor"
}

// DistributorColumns get sql column name.获取数据库列名
var DistributorColumns = struct {
	ID             string
	Code           string
	Name           string
	Email          string
	Mobile         string
	PointOfContact string
	Status         string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	Code:           "code",
	Name:           "name",
	Email:          "email",
	Mobile:         "mobile",
	PointOfContact: "point_of_contact",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

func (m *Distributor) FindByPrimaryKey(id string) (result Distributor, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Distributor) FindByLastInsertRuleRecode() (result Distributor, err error) {
	err = database.MysqlDB.Model(m).
		Where("code != ''").
		Order("SUBSTRING_INDEX(code, '-', -1) DESC").
		First(&result).Error
	return
}

func (m *Distributor) FindAll(whereCondition []database.WhereCondition) (results []Distributor, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}
