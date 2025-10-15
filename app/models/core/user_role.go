package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// UserRole [...]
type UserRole struct {
	ID        string    `gorm:"primaryKey;column:id" json:"-"`
	UserID    string    `gorm:"column:user_id" json:"userId"`
	User      User      `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	RoleID    string    `gorm:"column:role_id" json:"roleId"`
	Role      Role      `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID" json:"coreRoleList"`
	IsPrimary int8      `gorm:"column:is_primary" json:"isPrimary"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *UserRole) TableName() string {
	return "core_user_role"
}

// UserRoleColumns get sql column name.获取数据库列名
var UserRoleColumns = struct {
	ID        string
	UserID    string
	RoleID    string
	IsPrimary string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	RoleID:    "role_id",
	IsPrimary: "is_primary",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (m *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *UserRole) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *UserRole) Save() (result UserRole, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *UserRole) FindAll(whereCondition []database.WhereCondition) (results []UserRole, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *UserRole) FindByPrimaryKey(id string) (result UserRole, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *UserRole) FindByUserRole(userID string, roleID string) (result UserRole, err error) {
	err = database.MysqlDB.Model(m).
		Where("`user_id` = ?", userID).
		Where("`role_id` = ?", roleID).
		Find(&result).Error
	return
}

func (m *UserRole) FindByUser(userID string) (result UserRole, err error) {
	err = database.MysqlDB.Model(m).
		Where("`user_id` = ?", userID).
		Find(&result).Error
	return
}
