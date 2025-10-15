package entity

import "time"

// OutletLicense [...]
type OutletLicense struct {
	ID          string    `gorm:"primaryKey;column:id" json:"-"`
	LicenseType string    `gorm:"column:license_type" json:"licenseType"`
	LicenseNo   string    `gorm:"column:license_no" json:"licenseNo"`
	LicenseURL  string    `gorm:"column:license_url" json:"licenseUrl"`
	OutletID    string    `gorm:"column:outlet_id" json:"outletId"`
	Outlet      Outlet    `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *OutletLicense) TableName() string {
	return "ent_outlet_license"
}

// OutletLicenseColumns get sql column name.获取数据库列名
var OutletLicenseColumns = struct {
	ID          string
	LicenseType string
	LicenseNo   string
	LicenseURL  string
	OutletID    string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	LicenseType: "license_type",
	LicenseNo:   "license_no",
	LicenseURL:  "license_url",
	OutletID:    "outlet_id",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}
