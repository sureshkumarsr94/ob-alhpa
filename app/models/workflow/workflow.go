package workflow

import "time"

// Workflow [...]
type Workflow struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`
	Name         string    `gorm:"column:name" json:"name"`
	Description  string    `gorm:"column:description" json:"description"`
	StartDate    time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate      time.Time `gorm:"column:end_date" json:"endDate"`
	WorkflowType string    `gorm:"column:workflow_type" json:"workflowType"`
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Workflow) TableName() string {
	return "workflow"
}

// WorkflowColumns get sql column name.获取数据库列名
var WorkflowColumns = struct {
	ID           string
	Name         string
	Description  string
	StartDate    string
	EndDate      string
	WorkflowType string
	Status       string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Name:         "name",
	Description:  "description",
	StartDate:    "start_date",
	EndDate:      "end_date",
	WorkflowType: "workflow_type",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}
