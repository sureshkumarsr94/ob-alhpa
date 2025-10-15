package core

import "time"

// UserTerritory [...]
type UserTerritory struct {
	ID              string        `gorm:"primaryKey;column:id" json:"-"`
	UserID          string        `gorm:"column:user_id" json:"userId"`
	User            User          `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	TerritoryID     string        `gorm:"column:territory_id" json:"territoryId"`
	Territory       Territory     `gorm:"joinForeignKey:territory_id;foreignKey:id;references:TerritoryID" json:"coreTerritoryList"`
	TerritoryTypeID string        `gorm:"column:territory_type_id" json:"territoryTypeId"`
	TerritoryType   TerritoryType `gorm:"joinForeignKey:territory_type_id;foreignKey:id;references:TerritoryTypeID" json:"coreTerritoryTypeList"`
	Status          string        `gorm:"column:status" json:"status"`
	CreatedAt       time.Time     `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time     `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *UserTerritory) TableName() string {
	return "core_user_territory"
}

// UserTerritoryColumns get sql column name.获取数据库列名
var UserTerritoryColumns = struct {
	ID              string
	UserID          string
	TerritoryID     string
	TerritoryTypeID string
	Status          string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	UserID:          "user_id",
	TerritoryID:     "territory_id",
	TerritoryTypeID: "territory_type_id",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
