package sale

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// OrderScheme [...]
type OrderScheme struct {
	ID           string            `gorm:"primaryKey;column:id" json:"-"`
	OrderID      string            `gorm:"column:order_id" json:"orderId"`
	SaleOrder    Order             `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"saleOrderList"`
	SchemeID     string            `gorm:"column:scheme_id" json:"schemeId"`
	Scheme       coreModel.Scheme  `gorm:"joinForeignKey:scheme_id;foreignKey:id;references:SchemeID" json:"comSchemeList"`
	ProductID    string            `gorm:"column:product_id" json:"productId"`
	Product      coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	FreeSchemeID string            `gorm:"column:free_scheme_id" json:"freeSchemeId"`
	DiscQty      float64           `gorm:"column:disc_qty" json:"discQty"`
	DiscValue    float64           `gorm:"column:disc_value" json:"discValue"`
	Status       string            `gorm:"column:status" json:"status"`
	CreatedAt    time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderScheme) TableName() string {
	return "sale_order_scheme"
}

// OrderSchemeColumns get sql column name.获取数据库列名
var OrderSchemeColumns = struct {
	ID           string
	OrderID      string
	SchemeID     string
	ProductID    string
	FreeSchemeID string
	DiscQty      string
	DiscValue    string
	Status       string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	OrderID:      "order_id",
	SchemeID:     "scheme_id",
	ProductID:    "product_id",
	FreeSchemeID: "free_scheme_id",
	DiscQty:      "disc_qty",
	DiscValue:    "disc_value",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}
