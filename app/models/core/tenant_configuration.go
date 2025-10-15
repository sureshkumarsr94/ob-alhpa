package core

import "time"

// TenantConfiguration [...]
type TenantConfiguration struct {
	ID        string    `gorm:"primaryKey;column:id" json:"-"`
	Code      string    `gorm:"column:code" json:"code"`
	Data      string    `gorm:"column:data" json:"data"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.
func (m *TenantConfiguration) TableName() string {
	return "core_tenant_configuration"
}

// TenantConfigurationColumns get sql column name.
var TenantConfigurationColumns = struct {
	ID        string
	Code      string
	Data      string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Code:      "code",
	Data:      "data",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
