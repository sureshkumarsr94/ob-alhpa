package entity

import (
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// StockLog [...]
type StockLog struct {
	ID            string            `gorm:"primaryKey;column:id" json:"-"`
	Type          string            `gorm:"column:type" json:"type"`
	WarehouseID   string            `gorm:"column:warehouse_id" json:"warehouseId"`
	Warehouse     Warehouse         `gorm:"joinForeignKey:warehouse_id;foreignKey:id;references:WarehouseID" json:"entWarehouseList"`
	DistributorID string            `gorm:"column:distributor_id" json:"distributorId"`
	Distributor   Distributor       `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	UserID        string            `gorm:"column:user_id" json:"userId"`
	CoreUser      coreModel.User    `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	ProductID     string            `gorm:"column:product_id" json:"productId"`
	Product       coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	StockDate     time.Time         `gorm:"column:stock_date" json:"stockDate"`
	Qty           int               `gorm:"column:qty" json:"qty"`
	Description   string            `gorm:"column:description" json:"description"`
	BatchNo       string            `gorm:"column:batch_no" json:"batchNo"`
	Status        string            `gorm:"column:status" json:"status"`
	CreatedAt     time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *StockLog) TableName() string {
	return "ent_stock_log"
}

// StockLogColumns get sql column name.获取数据库列名
var StockLogColumns = struct {
	ID            string
	Type          string
	WarehouseID   string
	DistributorID string
	UserID        string
	ProductID     string
	StockDate     string
	Qty           string
	Description   string
	BatchNo       string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Type:          "type",
	WarehouseID:   "warehouse_id",
	DistributorID: "distributor_id",
	UserID:        "user_id",
	ProductID:     "product_id",
	StockDate:     "stock_date",
	Qty:           "qty",
	Description:   "description",
	BatchNo:       "batch_no",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
