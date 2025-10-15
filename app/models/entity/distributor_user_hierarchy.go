package entity

import "time"

// DistributorUserHierarchy [...]
type DistributorUserHierarchy struct {
	ID                string          `gorm:"primaryKey;column:id" json:"-"`
	DistributorID     string          `gorm:"column:distributor_id" json:"distributorId"`
	Distributor       Distributor     `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	DistributorUserID string          `gorm:"column:distributor_user_id" json:"distributorUserId"`
	DistributorUser   DistributorUser `gorm:"joinForeignKey:distributor_user_id;foreignKey:id;references:DistributorUserID" json:"entDistributorUserList"`
	SupervisorID      string          `gorm:"column:supervisor_id" json:"supervisorId"`
	Status            string          `gorm:"column:status" json:"status"`
	CreatedAt         time.Time       `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time       `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *DistributorUserHierarchy) TableName() string {
	return "ent_distributor_user_hierarchy"
}

// DistributorUserHierarchyColumns get sql column name.获取数据库列名
var DistributorUserHierarchyColumns = struct {
	ID                string
	DistributorID     string
	DistributorUserID string
	SupervisorID      string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	DistributorID:     "distributor_id",
	DistributorUserID: "distributor_user_id",
	SupervisorID:      "supervisor_id",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}
