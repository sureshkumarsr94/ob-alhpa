package entity

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// CategoryUserScope [...]
type CategoryUserScope struct {
	ID         string             `gorm:"primaryKey;column:id" json:"-"`
	CategoryID string             `gorm:"column:category_id" json:"categoryId"`
	Category   coreModel.Category `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryID" json:"comCategoryList"`
	UserID     string             `gorm:"column:user_id" json:"userId"`
	User       coreModel.User     `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	Status     string             `gorm:"column:status" json:"status"`
	CreatedAt  time.Time          `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time          `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *CategoryUserScope) TableName() string {
	return "ent_category_user_scope"
}

// CategoryUserScopeColumns get sql column name.获取数据库列名
var CategoryUserScopeColumns = struct {
	ID         string
	CategoryID string
	UserID     string
	Status     string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	CategoryID: "category_id",
	UserID:     "user_id",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}
