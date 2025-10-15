package workflow

import (
	"gorm.io/datatypes"
	"time"
)

// Stage [...]
type Stage struct {
	ID            string         `gorm:"primaryKey;column:id" json:"-"`
	WorkflowID    string         `gorm:"column:workflow_id" json:"workflowId"`
	Workflow      Workflow       `gorm:"joinForeignKey:workflow_id;foreignKey:id;references:WorkflowID" json:"workflowList"`
	Name          string         `gorm:"column:name" json:"name"`
	Description   string         `gorm:"column:description" json:"description"`
	Sequence      string         `gorm:"column:sequence" json:"sequence"`
	Configuration datatypes.JSON `gorm:"column:configuration" json:"configuration"`
	Status        string         `gorm:"column:status" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Stage) TableName() string {
	return "workflow_stage"
}

// StageColumns get sql column name.获取数据库列名
var StageColumns = struct {
	ID            string
	WorkflowID    string
	Name          string
	Description   string
	Sequence      string
	Configuration string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	WorkflowID:    "workflow_id",
	Name:          "name",
	Description:   "description",
	Sequence:      "sequence",
	Configuration: "configuration",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
