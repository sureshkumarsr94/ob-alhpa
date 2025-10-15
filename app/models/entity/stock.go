package entity

import (
	"github.com/guregu/null"
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// Stock [...]
type Stock struct {
	ID            string            `gorm:"primaryKey;column:id" json:"-"`
	WarehouseID   string            `gorm:"column:warehouse_id" json:"warehouseId"`
	Warehouse     Warehouse         `gorm:"joinForeignKey:warehouse_id;foreignKey:id;references:WarehouseID" json:"entWarehouseList"`
	DistributorID string            `gorm:"column:distributor_id" json:"distributorId"`
	Distributor   Distributor       `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	UserID        null.String       `gorm:"column:user_id" json:"userId"`
	CoreUser      coreModel.User    `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	ProductID     string            `gorm:"column:product_id" json:"productId"`
	ComProduct    coreModel.Product `gorm:"joinForeignKey:product_id;foreignKey:id;references:ProductID" json:"comProductList"`
	Qty           int               `gorm:"column:qty" json:"qty"`
	Description   string            `gorm:"column:description" json:"description"`
	BatchNo       string            `gorm:"column:batch_no" json:"batchNo"`
	Status        string            `gorm:"column:status" json:"status"`
	CreatedAt     time.Time         `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Stock) TableName() string {
	return "ent_stock"
}

// StockColumns get sql column name.获取数据库列名
var StockColumns = struct {
	ID            string
	WarehouseID   string
	DistributorID string
	UserID        string
	ProductID     string
	Qty           string
	Description   string
	BatchNo       string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	WarehouseID:   "warehouse_id",
	DistributorID: "distributor_id",
	UserID:        "user_id",
	ProductID:     "product_id",
	Qty:           "qty",
	Description:   "description",
	BatchNo:       "batch_no",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
