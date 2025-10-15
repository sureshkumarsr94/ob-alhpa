package sale

import (
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"infopack.co.in/offybox/app/models/entity"
	"time"
)

// Order [...]
type Order struct {
	ID              string               `gorm:"primaryKey;column:id" json:"-"`
	OutletID        string               `gorm:"column:outlet_id" json:"outletId"`
	Outlet          entity.Outlet        `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	Type            string               `gorm:"column:type" json:"type"`
	ItemCount       int                  `gorm:"column:item_count" json:"itemCount"`
	NetAmount       float64              `gorm:"column:net_amount" json:"netAmount"`
	DiscountAmount  float64              `gorm:"column:discount_amount" json:"discountAmount"`
	TaxAmount       float64              `gorm:"column:tax_amount" json:"taxAmount"`
	TotalAmount     float64              `gorm:"column:total_amount" json:"totalAmount"`
	QuotationDate   time.Time            `gorm:"column:quotation_date" json:"quotationDate"`
	CreatedBy       string               `gorm:"column:created_by" json:"createdBy"`
	CoreUser        coreModel.User       `gorm:"joinForeignKey:created_by;foreignKey:id;references:CreatedBy" json:"coreUserList"`
	Remarks         string               `gorm:"column:remarks" json:"remarks"`
	Latitude        string               `gorm:"column:latitude" json:"latitude"`
	Longitude       string               `gorm:"column:longitude" json:"longitude"`
	OutletAddressID string               `gorm:"column:outlet_address_id" json:"outletAddressId"`
	OutletAddress   entity.OutletAddress `gorm:"joinForeignKey:outlet_address_id;foreignKey:id;references:OutletAddressID" json:"entOutletAddressList"`
	FileURL         string               `gorm:"column:file_url" json:"fileUrl"`
	DistributorID   string               `gorm:"column:distributor_id" json:"distributorId"`
	Distributor     entity.Distributor   `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	OrderStatus     string               `gorm:"column:order_status" json:"orderStatus"`
	Status          string               `gorm:"column:status" json:"status"`
	CreatedAt       time.Time            `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time            `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Order) TableName() string {
	return "sale_order"
}

// OrderColumns get sql column name.获取数据库列名
var OrderColumns = struct {
	ID              string
	OutletID        string
	Type            string
	ItemCount       string
	NetAmount       string
	DiscountAmount  string
	TaxAmount       string
	TotalAmount     string
	QuotationDate   string
	CreatedBy       string
	Remarks         string
	Latitude        string
	Longitude       string
	OutletAddressID string
	FileURL         string
	DistributorID   string
	OrderStatus     string
	Status          string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	OutletID:        "outlet_id",
	Type:            "type",
	ItemCount:       "item_count",
	NetAmount:       "net_amount",
	DiscountAmount:  "discount_amount",
	TaxAmount:       "tax_amount",
	TotalAmount:     "total_amount",
	QuotationDate:   "quotation_date",
	CreatedBy:       "created_by",
	Remarks:         "remarks",
	Latitude:        "latitude",
	Longitude:       "longitude",
	OutletAddressID: "outlet_address_id",
	FileURL:         "file_url",
	DistributorID:   "distributor_id",
	OrderStatus:     "order_status",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

func (m *Order) BeforeCreate(tx *gorm.DB) (err error) {
	m.Status = "ACTIVE"
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Order) FindByPrimaryKey(id string) (result Order, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}
