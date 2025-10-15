package sale

import (
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// OrderItem [...]
type OrderItem struct {
	ID             string            `gorm:"primaryKey;column:id" json:"-"`
	ProductID      string            `gorm:"column:product_id" json:"productId"`
	Product        coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	FreeProductID  null.String       `gorm:"column:free_product_id" json:"freeProductId"`
	Qty            int               `gorm:"column:qty" json:"qty"`
	FreeQty        int               `gorm:"column:free_qty" json:"freeQty"`
	Price          float64           `gorm:"column:price" json:"price"`
	NetAmount      float64           `gorm:"column:net_amount" json:"netAmount"`
	DiscountAmount null.Float        `gorm:"column:discount_amount" json:"discountAmount"`
	TaxAmount      float64           `gorm:"column:tax_amount" json:"taxAmount"`
	TotalAmount    float64           `gorm:"column:total_amount" json:"totalAmount"`
	TaxPercentage  float64           `gorm:"column:tax_percentage" json:"taxPercentage"`
	SchemeID       null.Float        `gorm:"column:scheme_id" json:"schemeId"`
	Scheme         coreModel.Scheme  `gorm:"joinForeignKey:scheme_id;foreignKey:id;references:SchemeID" json:"comSchemeList"`
	OrderID        string            `gorm:"column:order_id" json:"orderId"`
	Order          Order             `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"saleOrderList"`
	Status         string            `gorm:"column:status" json:"status"`
	CreatedAt      time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderItem) TableName() string {
	return "sale_order_item"
}

// OrderItemColumns get sql column name.获取数据库列名
var OrderItemColumns = struct {
	ID             string
	ProductID      string
	FreeProductID  string
	Qty            string
	FreeQty        string
	Price          string
	NetAmount      string
	DiscountAmount string
	TaxAmount      string
	TotalAmount    string
	TaxPercentage  string
	SchemeID       string
	OrderID        string
	Status         string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	ProductID:      "product_id",
	FreeProductID:  "free_product_id",
	Qty:            "qty",
	FreeQty:        "free_qty",
	Price:          "price",
	NetAmount:      "net_amount",
	DiscountAmount: "discount_amount",
	TaxAmount:      "tax_amount",
	TotalAmount:    "total_amount",
	TaxPercentage:  "tax_percentage",
	SchemeID:       "scheme_id",
	OrderID:        "order_id",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

func (m *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *OrderItem) FindByPrimaryKey(id string) (result OrderItem, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}
