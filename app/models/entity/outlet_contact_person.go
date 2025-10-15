package entity

import "time"

// OutletContactPerson [...]
type OutletContactPerson struct {
	ID        string    `gorm:"primaryKey;column:id" json:"-"`
	Title     string    `gorm:"column:title" json:"title"`
	Name      string    `gorm:"column:name" json:"name"`
	Mobile    string    `gorm:"column:mobile" json:"mobile"`
	Email     string    `gorm:"column:email" json:"email"`
	Position  string    `gorm:"column:position" json:"position"`
	OutletID  string    `gorm:"column:outlet_id" json:"outletId"`
	Outlet    Outlet    `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *OutletContactPerson) TableName() string {
	return "ent_outlet_contact_person"
}

// OutletContactPersonColumns get sql column name.获取数据库列名
var OutletContactPersonColumns = struct {
	ID        string
	Title     string
	Name      string
	Mobile    string
	Email     string
	Position  string
	OutletID  string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Title:     "title",
	Name:      "name",
	Mobile:    "mobile",
	Email:     "email",
	Position:  "position",
	OutletID:  "outlet_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
