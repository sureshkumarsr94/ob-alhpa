package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Role [...]
type Role struct {
	ID          string      `gorm:"primaryKey;column:id" json:"-"`
	UserType    string      `gorm:"column:user_type" json:"userType"`
	Code        string      `gorm:"column:code" json:"code"`
	Name        string      `gorm:"column:name" json:"name"`
	Description string      `gorm:"column:description" json:"description"`
	RoleID      null.String `gorm:"column:role_id" json:"roleId"`
	//ParentRole    *Role       `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID" json:"coreRoleList"`
	DataAccess    string      `gorm:"column:data_access" json:"dataAccess"`
	DistributorID null.String `gorm:"column:distributor_id" json:"distributorId"`
	Status        string      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time   `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Role) TableName() string {
	return "core_role"
}

// RoleColumns get sql column name.获取数据库列名
var RoleColumns = struct {
	ID            string
	UserType      string
	Code          string
	Name          string
	Description   string
	RoleID        string
	DataAccess    string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	UserType:      "user_type",
	Code:          "code",
	Name:          "name",
	Description:   "description",
	RoleID:        "role_id",
	DataAccess:    "data_access",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *Role) FindByPrimaryKey(id string) (result Role, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Role) FindByDistributorAndRole(id string, role string) (result Role, err error) {
	err = database.MysqlDB.Model(m).
		Where("`distributor_id` = ? AND `code` = ?", id, role).Find(&result).Error
	return
}

func (m *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Role) FindAll(whereCondition []database.WhereCondition) (results []Role, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Role) Save() (result Role, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
