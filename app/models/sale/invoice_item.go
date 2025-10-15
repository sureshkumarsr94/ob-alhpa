package sale

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// InvoiceItem [...]
type InvoiceItem struct {
	ID             string            `gorm:"primaryKey;column:id" json:"-"`
	ProductID      string            `gorm:"column:product_id" json:"productId"`
	Product        coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	FreeProductID  string            `gorm:"column:free_product_id" json:"freeProductId"`
	Qty            int               `gorm:"column:qty" json:"qty"`
	Price          float64           `gorm:"column:price" json:"price"`
	NetAmount      float64           `gorm:"column:net_amount" json:"netAmount"`
	DiscountAmount float64           `gorm:"column:discount_amount" json:"discountAmount"`
	TaxAmount      float64           `gorm:"column:tax_amount" json:"taxAmount"`
	TotalAmount    float64           `gorm:"column:total_amount" json:"totalAmount"`
	TaxPercentage  float64           `gorm:"column:tax_percentage" json:"taxPercentage"`
	SchemeID       string            `gorm:"column:scheme_id" json:"schemeId"`
	Scheme         coreModel.Scheme  `gorm:"joinForeignKey:scheme_id;foreignKey:id;references:SchemeID" json:"comSchemeList"`
	InvoiceID      string            `gorm:"column:invoice_id" json:"invoiceId"`
	Invoice        Invoice           `gorm:"joinForeignKey:invoice_id;foreignKey:id;references:InvoiceID" json:"saleInvoiceList"`
	Status         string            `gorm:"column:status" json:"status"`
	CreatedAt      time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *InvoiceItem) TableName() string {
	return "sale_invoice_item"
}

// InvoiceItemColumns get sql column name.获取数据库列名
var InvoiceItemColumns = struct {
	ID             string
	ProductID      string
	FreeProductID  string
	Qty            string
	Price          string
	NetAmount      string
	DiscountAmount string
	TaxAmount      string
	TotalAmount    string
	TaxPercentage  string
	SchemeID       string
	InvoiceID      string
	Status         string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	ProductID:      "product_id",
	FreeProductID:  "free_product_id",
	Qty:            "qty",
	Price:          "price",
	NetAmount:      "net_amount",
	DiscountAmount: "discount_amount",
	TaxAmount:      "tax_amount",
	TotalAmount:    "total_amount",
	TaxPercentage:  "tax_percentage",
	SchemeID:       "scheme_id",
	InvoiceID:      "invoice_id",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}
