package workflow

import (
	"gorm.io/datatypes"
	"time"
)

// Instance [...]
type Instance struct {
	ID                       string         `gorm:"primaryKey;column:id" json:"-"`
	WorkflowID               string         `gorm:"column:workflow_id" json:"workflowId"`
	Workflow                 Workflow       `gorm:"joinForeignKey:workflow_id;foreignKey:id;references:WorkflowID" json:"workflowList"`
	SourceID                 string         `gorm:"column:source_id" json:"sourceId"`
	SourceType               string         `gorm:"column:source_type" json:"sourceType"`
	ReferenceID              string         `gorm:"column:reference_id" json:"referenceId"`
	ReferenceType            string         `gorm:"column:reference_type" json:"referenceType"`
	Definition               datatypes.JSON `gorm:"column:definition" json:"definition"`
	PreviousInstanceID       string         `gorm:"column:previous_instance_id" json:"previousInstanceId"`
	PreviousWorkflowInstance *Instance      `gorm:"joinForeignKey:previous_instance_id;foreignKey:id;references:PreviousInstanceID" json:"workflowInstanceList"`
	ParentInstanceID         string         `gorm:"column:parent_instance_id" json:"parentInstanceId"`
	StartDate                time.Time      `gorm:"column:start_date" json:"startDate"`
	EndDate                  time.Time      `gorm:"column:end_date" json:"endDate"`
	ActiveTaskID             string         `gorm:"column:active_task_id" json:"activeTaskId"`
	ActiveTask               *Task          `gorm:"joinForeignKey:active_task_id;foreignKey:id;references:ActiveTaskID" json:"activeTask"`
	CompletedTaskID          string         `gorm:"column:completed_task_id" json:"completedTaskId"`
	CompleteTask             *Task          `gorm:"joinForeignKey:completed_task_id;foreignKey:id;references:CompletedTaskID" json:"completedTask"`
	Status                   string         `gorm:"column:status" json:"status"`
	CreatedAt                time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt                time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.
func (m *Instance) TableName() string {
	return "workflow_instance"
}

// InstanceColumns get sql column name.
var InstanceColumns = struct {
	ID                 string
	WorkflowID         string
	SourceID           string
	SourceType         string
	ReferenceID        string
	ReferenceType      string
	Definition         string
	PreviousInstanceID string
	ParentInstanceID   string
	StartDate          string
	EndDate            string
	ActiveTaskID       string
	CompletedTaskID    string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	WorkflowID:         "workflow_id",
	SourceID:           "source_id",
	SourceType:         "source_type",
	ReferenceID:        "reference_id",
	ReferenceType:      "reference_type",
	Definition:         "definition",
	PreviousInstanceID: "previous_instance_id",
	ParentInstanceID:   "parent_instance_id",
	StartDate:          "start_date",
	EndDate:            "end_date",
	ActiveTaskID:       "active_task_id",
	CompletedTaskID:    "completed_task_id",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}
