package sale

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"infopack.co.in/offybox/app/models/entity"
	"infopack.co.in/offybox/app/models/workflow"
	"time"
)

// Participant [...]
type Participant struct {
	ID              string             `gorm:"primaryKey;column:id" json:"-"`
	OrderID         string             `gorm:"column:order_id" json:"orderId"`
	SaleOrder       Order              `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"saleOrderList"`
	UserID          string             `gorm:"column:user_id" json:"userId"`
	CoreUser        coreModel.User     `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	DistributorID   string             `gorm:"column:distributor_id" json:"distributorId"`
	Distributor     entity.Distributor `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	ParticipantType string             `gorm:"column:participant_type" json:"participantType"`
	TaskID          string             `gorm:"column:task_id" json:"taskId"`
	WorkflowTask    workflow.Task      `gorm:"joinForeignKey:task_id;foreignKey:id;references:TaskID" json:"workflowTaskList"`
	Status          string             `gorm:"column:status" json:"status"`
	CreatedAt       time.Time          `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time          `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Participant) TableName() string {
	return "sale_participant"
}

// ParticipantColumns get sql column name.获取数据库列名
var ParticipantColumns = struct {
	ID              string
	OrderID         string
	UserID          string
	DistributorID   string
	ParticipantType string
	TaskID          string
	Status          string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	OrderID:         "order_id",
	UserID:          "user_id",
	DistributorID:   "distributor_id",
	ParticipantType: "participant_type",
	TaskID:          "task_id",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
