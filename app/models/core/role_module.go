package core

import (
	"gorm.io/datatypes"
	"time"
)

// RoleModule [...]
type RoleModule struct {
	ID                string         `gorm:"primaryKey;column:id" json:"-"`
	RoleID            string         `gorm:"column:role_id" json:"roleId"`
	Role              Role           `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID" json:"coreRoleList"`
	ModuleID          string         `gorm:"column:module_id" json:"moduleId"`
	Module            Module         `gorm:"joinForeignKey:module_id;foreignKey:id;references:ModuleID" json:"coreModuleList"`
	AllowedPermission datatypes.JSON `gorm:"column:allowed_permission" json:"allowedPermission"`
	DataAccess        string         `gorm:"column:data_access" json:"dataAccess"`
	Status            string         `gorm:"column:status" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *RoleModule) TableName() string {
	return "core_role_module"
}

// RoleModuleColumns get sql column name.获取数据库列名
var RoleModuleColumns = struct {
	ID                string
	RoleID            string
	ModuleID          string
	AllowedPermission string
	DataAccess        string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	RoleID:            "role_id",
	ModuleID:          "module_id",
	AllowedPermission: "allowed_permission",
	DataAccess:        "data_access",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}
