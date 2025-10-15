package sale

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"infopack.co.in/offybox/app/models/entity"
	"time"
)

// Trip [...]
type Trip struct {
	ID            string             `gorm:"primaryKey;column:id" json:"-"`
	UserID        string             `gorm:"column:user_id" json:"userId"`
	CoreUser      coreModel.User     `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	StartDate     time.Time          `gorm:"column:start_date" json:"startDate"`
	EndDate       null.Time          `gorm:"column:end_date" json:"endDate"`
	VehicleNumber string             `gorm:"column:vehicle_number" json:"vehicleNumber"`
	VehicleType   string             `gorm:"column:vehicle_type" json:"vehicleType"`
	VehicleName   string             `gorm:"column:vehicle_name" json:"vehicleName"`
	DriverName    string             `gorm:"column:driver_name" json:"driverName"`
	DriverProof   string             `gorm:"column:driver_proof" json:"driverProof"`
	StartKm       float64            `gorm:"column:start_km" json:"startKm"`
	EndKm         float64            `gorm:"column:end_km" json:"endKm"`
	LoadedQty     float64            `gorm:"column:loaded_qty" json:"loadedQty"`
	ReturnedQty   float64            `gorm:"column:returned_qty" json:"returnedQty"`
	DamagedQty    float64            `gorm:"column:damaged_qty" json:"damagedQty"`
	DistributorID null.String        `gorm:"column:distributor_id" json:"distributorId"`
	Distributor   entity.Distributor `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	Status        string             `gorm:"column:status" json:"status"`
	CreatedAt     time.Time          `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time          `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Trip) TableName() string {
	return "sale_trip"
}

// TripColumns get sql column name.获取数据库列名
var TripColumns = struct {
	ID            string
	UserID        string
	StartDate     string
	EndDate       string
	VehicleNumber string
	VehicleType   string
	VehicleName   string
	DriverName    string
	DriverProof   string
	StartKm       string
	EndKm         string
	LoadedQty     string
	ReturnedQty   string
	DamagedQty    string
	DistributorID string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	UserID:        "user_id",
	StartDate:     "start_date",
	EndDate:       "end_date",
	VehicleNumber: "vehicle_number",
	VehicleType:   "vehicle_type",
	VehicleName:   "vehicle_name",
	DriverName:    "driver_name",
	DriverProof:   "driver_proof",
	StartKm:       "start_km",
	EndKm:         "end_km",
	LoadedQty:     "loaded_qty",
	ReturnedQty:   "returned_qty",
	DamagedQty:    "damaged_qty",
	DistributorID: "distributor_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *Trip) FindByPrimaryKey(id string) (result Trip, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Trip) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Trip) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Trip) FindAll(whereCondition []database.WhereCondition) (results []Trip, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Trip) Save() (result Trip, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
