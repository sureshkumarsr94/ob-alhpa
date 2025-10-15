package core

import "time"

// LookupMaster [...]
type LookupMaster struct {
	ID            string    `gorm:"primaryKey;column:id" json:"-"`
	LuKey         string    `gorm:"column:lu_key" json:"luKey"`
	LuName        string    `gorm:"column:lu_name" json:"luName"`
	LuValue       string    `gorm:"column:lu_value" json:"luValue"`
	GroupCode     string    `gorm:"column:group_code" json:"groupCode"`
	DistributorID string    `gorm:"column:distributor_id" json:"distributorId"`
	Status        string    `gorm:"column:status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *LookupMaster) TableName() string {
	return "core_lookup_master"
}

// LookupMasterColumns get sql column name.获取数据库列名
var LookupMasterColumns = struct {
	ID            string
	LuKey         string
	LuName        string
	LuValue       string
	GroupCode     string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	LuKey:         "lu_key",
	LuName:        "lu_name",
	LuValue:       "lu_value",
	GroupCode:     "group_code",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
