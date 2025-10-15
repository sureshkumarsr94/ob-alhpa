package workflow

import (
	"gorm.io/datatypes"
	"time"
)

// Step [...]
type Step struct {
	ID                 string         `gorm:"primaryKey;column:id" json:"-"`
	WorkflowID         string         `gorm:"column:workflow_id" json:"workflowId"`
	Workflow           Workflow       `gorm:"joinForeignKey:workflow_id;foreignKey:id;references:WorkflowID" json:"workflowList"`
	WorkflowStageID    string         `gorm:"column:workflow_stage_id" json:"workflowStageId"`
	WorkflowStage      Stage          `gorm:"joinForeignKey:workflow_stage_id;foreignKey:id;references:WorkflowStageID" json:"workflowStageList"`
	Name               string         `gorm:"column:name" json:"name"`
	Description        string         `gorm:"column:description" json:"description"`
	StepType           string         `gorm:"column:step_type" json:"stepType"`
	Configuration      datatypes.JSON `gorm:"column:configuration" json:"configuration"`
	UIComponent        string         `gorm:"column:ui_component" json:"uiComponent"`
	AutomaticComponent string         `gorm:"column:automatic_component" json:"automaticComponent"`
	AllocationRuleID   string         `gorm:"column:allocation_rule_id" json:"allocationRuleId"`
	WorkflowRule       Rule           `gorm:"joinForeignKey:allocation_rule_id;foreignKey:id;references:AllocationRuleID" json:"workflowRuleList"`
	DisplayRuleID      string         `gorm:"column:display_rule_id" json:"displayRuleId"`
	Status             string         `gorm:"column:status" json:"status"`
	CreatedAt          time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Step) TableName() string {
	return "workflow_step"
}

// StepColumns get sql column name.获取数据库列名
var StepColumns = struct {
	ID                 string
	WorkflowID         string
	WorkflowStageID    string
	Name               string
	Description        string
	StepType           string
	Configuration      string
	UIComponent        string
	AutomaticComponent string
	AllocationRuleID   string
	DisplayRuleID      string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	WorkflowID:         "workflow_id",
	WorkflowStageID:    "workflow_stage_id",
	Name:               "name",
	Description:        "description",
	StepType:           "step_type",
	Configuration:      "configuration",
	UIComponent:        "ui_component",
	AutomaticComponent: "automatic_component",
	AllocationRuleID:   "allocation_rule_id",
	DisplayRuleID:      "display_rule_id",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}
