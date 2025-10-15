package sale

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"infopack.co.in/offybox/app/models/entity"
	"time"
)

// Invoice [...]
type Invoice struct {
	ID              string               `gorm:"primaryKey;column:id" json:"-"`
	OutletID        string               `gorm:"column:outlet_id" json:"outletId"`
	Outlet          entity.Outlet        `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
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
	OrderID         string               `gorm:"column:order_id" json:"orderId"`
	Order           Order                `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"saleOrderList"`
	PaidAmount      float64              `gorm:"column:paid_amount" json:"paidAmount"`
	UnpaidAmount    float64              `gorm:"column:unpaid_amount" json:"unpaidAmount"`
	DistributorID   string               `gorm:"column:distributor_id" json:"distributorId"`
	Distributor     entity.Distributor   `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	Status          string               `gorm:"column:status" json:"status"`
	CreatedAt       time.Time            `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time            `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Invoice) TableName() string {
	return "sale_invoice"
}

// InvoiceColumns get sql column name.获取数据库列名
var InvoiceColumns = struct {
	ID              string
	OutletID        string
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
	OrderID         string
	PaidAmount      string
	UnpaidAmount    string
	DistributorID   string
	Status          string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	OutletID:        "outlet_id",
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
	OrderID:         "order_id",
	PaidAmount:      "paid_amount",
	UnpaidAmount:    "unpaid_amount",
	DistributorID:   "distributor_id",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

func (m *Invoice) FindByPrimaryKey(id string) (result Invoice, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Invoice) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Invoice) FindAll(whereCondition []database.WhereCondition) (results []Invoice, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Invoice) Save() (result Invoice, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
