package workflow

import (
	"gorm.io/datatypes"
	"time"
)

// Rule [...]
type Rule struct {
	ID        string         `gorm:"primaryKey;column:id" json:"-"`
	Type      string         `gorm:"column:type" json:"type"`
	Rule      datatypes.JSON `gorm:"column:rule" json:"rule"`
	Name      string         `gorm:"column:name" json:"name"`
	Status    string         `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Rule) TableName() string {
	return "workflow_rule"
}

// RuleColumns get sql column name.获取数据库列名
var RuleColumns = struct {
	ID        string
	Type      string
	Rule      string
	Name      string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Type:      "type",
	Rule:      "rule",
	Name:      "name",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
