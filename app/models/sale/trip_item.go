package sale

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"infopack.co.in/offybox/app/models/entity"
	"time"
)

// TripItem [...]
type TripItem struct {
	ID        string            `gorm:"primaryKey;column:id" json:"-"`
	TripID    string            `gorm:"column:TripItem_id" json:"TripItemId"`
	Trip      Trip              `gorm:"joinForeignKey:TripItem_id;foreignKey:id;references:TripID" json:"saleTripItemList"`
	OrderID   string            `gorm:"column:order_id" json:"orderId"`
	Order     Order             `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"saleOrderList"`
	InvoiceID string            `gorm:"column:invoice_id" json:"invoiceId"`
	Invoice   Invoice           `gorm:"joinForeignKey:invoice_id;foreignKey:id;references:InvoiceID" json:"saleInvoiceList"`
	OutletID  string            `gorm:"column:outlet_id" json:"outletId"`
	Outlet    entity.Outlet     `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	ProductID string            `gorm:"column:product_id" json:"productId"`
	Product   coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	Qty       float64           `gorm:"column:qty" json:"qty"`
	FreeQty   float64           `gorm:"column:free_qty" json:"freeQty"`
	Remarks   string            `gorm:"column:remarks" json:"remarks"`
	Status    string            `gorm:"column:status" json:"status"`
	CreatedAt time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *TripItem) TableName() string {
	return "sale_TripItem_item"
}

// TripItemColumns get sql column name.获取数据库列名
var TripItemColumns = struct {
	ID         string
	TripItemID string
	OrderID    string
	InvoiceID  string
	OutletID   string
	ProductID  string
	Qty        string
	FreeQty    string
	Remarks    string
	Status     string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	TripItemID: "TripItem_id",
	OrderID:    "order_id",
	InvoiceID:  "invoice_id",
	OutletID:   "outlet_id",
	ProductID:  "product_id",
	Qty:        "qty",
	FreeQty:    "free_qty",
	Remarks:    "remarks",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

func (m *TripItem) FindByPrimaryKey(id string) (result TripItem, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *TripItem) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *TripItem) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *TripItem) FindAll(whereCondition []database.WhereCondition) (results []TripItem, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *TripItem) Save() (result TripItem, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *TripItem) FindByTripID(id string) (result []TripItem, err error) {
	err = database.MysqlDB.Model(m).
		Where("`trip_id` = ?", id).
		Where("`status` = ?", coreModel.StatusActive).
		Find(&result).Error
	return
}
