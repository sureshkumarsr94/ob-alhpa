package core

import (
	"gorm.io/datatypes"
	"time"
)

// Module [...]
type Module struct {
	ID                string         `gorm:"primaryKey;column:id" json:"-"`
	UserType          string         `gorm:"column:user_type" json:"userType"`
	System            string         `gorm:"column:system" json:"system"`
	Code              string         `gorm:"column:code" json:"code"`
	ParentModuleID    string         `gorm:"column:parent_module_id" json:"parentModuleId"`
	ParentModule      *Module        `gorm:"joinForeignKey:parent_module_id;foreignKey:id;references:ParentModuleID" json:"coreModuleList"`
	Name              string         `gorm:"column:name" json:"name"`
	Description       string         `gorm:"column:description" json:"description"`
	AllowedPermission datatypes.JSON `gorm:"column:allowed_permission" json:"allowedPermission"`
	URL               string         `gorm:"column:url" json:"url"`
	Icon              string         `gorm:"column:icon" json:"icon"`
	Sequence          int            `gorm:"column:sequence" json:"sequence"`
	Target            string         `gorm:"column:target" json:"target"`
	Status            string         `gorm:"column:status" json:"status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Module) TableName() string {
	return "core_module"
}

// ModuleColumns get sql column name.获取数据库列名
var ModuleColumns = struct {
	ID                string
	UserType          string
	System            string
	Code              string
	ParentModuleID    string
	Name              string
	Description       string
	AllowedPermission string
	URL               string
	Icon              string
	Sequence          string
	Target            string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	UserType:          "user_type",
	System:            "system",
	Code:              "code",
	ParentModuleID:    "parent_module_id",
	Name:              "name",
	Description:       "description",
	AllowedPermission: "allowed_permission",
	URL:               "url",
	Icon:              "icon",
	Sequence:          "sequence",
	Target:            "target",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}
