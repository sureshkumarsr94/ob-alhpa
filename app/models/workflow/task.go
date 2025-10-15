package workflow

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// Task [...]
type Task struct {
	ID                 string             `gorm:"primaryKey;column:id" json:"-"`
	WorkflowInstanceID string             `gorm:"column:workflow_instance_id" json:"workflowInstanceId"`
	WorkflowInstance   Instance           `gorm:"joinForeignKey:workflow_instance_id;foreignKey:id;references:WorkflowInstanceID" json:"workflowInstanceList"`
	WorkflowID         string             `gorm:"column:workflow_id" json:"workflowId"`
	Workflow           Workflow           `gorm:"joinForeignKey:workflow_id;foreignKey:id;references:WorkflowID" json:"workflowList"`
	WorkflowStepID     string             `gorm:"column:workflow_step_id" json:"workflowStepId"`
	WorkflowStep       Step               `gorm:"joinForeignKey:workflow_step_id;foreignKey:id;references:WorkflowStepID" json:"workflowStepList"`
	UserID             string             `gorm:"column:user_id" json:"userId"`
	CoreUser           coreModel.User     `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	UserRoleID         string             `gorm:"column:user_role_id" json:"userRoleId"`
	CoreUserRole       coreModel.UserRole `gorm:"joinForeignKey:user_role_id;foreignKey:id;references:UserRoleID" json:"coreUserRoleList"`
	AssignedDate       time.Time          `gorm:"column:assigned_date" json:"assignedDate"`
	ClosedDate         time.Time          `gorm:"column:closed_date" json:"closedDate"`
	Status             string             `gorm:"column:status" json:"status"`
	CreatedAt          time.Time          `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt          time.Time          `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "workflow_task"
}

// TaskColumns get sql column name.获取数据库列名
var TaskColumns = struct {
	ID                 string
	WorkflowInstanceID string
	WorkflowID         string
	WorkflowStepID     string
	UserID             string
	UserRoleID         string
	AssignedDate       string
	ClosedDate         string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	WorkflowInstanceID: "workflow_instance_id",
	WorkflowID:         "workflow_id",
	WorkflowStepID:     "workflow_step_id",
	UserID:             "user_id",
	UserRoleID:         "user_role_id",
	AssignedDate:       "assigned_date",
	ClosedDate:         "closed_date",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}
