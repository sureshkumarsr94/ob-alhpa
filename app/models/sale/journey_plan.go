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

// JourneyPlan [...]
type JourneyPlan struct {
	ID           string         `gorm:"primaryKey;column:id" json:"-"`
	OutletID     string         `gorm:"column:outlet_id" json:"outletId"`
	Outlet       entity.Outlet  `gorm:"joinForeignKey:outlet_id;foreignKey:id;references:OutletID" json:"entOutletList"`
	UserID       string         `gorm:"column:user_id" json:"userId"`
	CoreUser     coreModel.User `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	AssignedDate time.Time      `gorm:"column:assigned_date" json:"assignedDate"`
	ClosedDate   null.Time      `gorm:"column:closed_date" json:"closedDate"`
	Remarks      string         `gorm:"column:remarks" json:"remarks"`
	Status       string         `gorm:"column:status" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *JourneyPlan) TableName() string {
	return "sale_journey_plan"
}

// JourneyPlanColumns get sql column name.获取数据库列名
var JourneyPlanColumns = struct {
	ID           string
	OutletID     string
	UserID       string
	AssignedDate string
	ClosedDate   string
	Remarks      string
	Status       string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	OutletID:     "outlet_id",
	UserID:       "user_id",
	AssignedDate: "assigned_date",
	ClosedDate:   "closed_date",
	Remarks:      "remarks",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

func (m *JourneyPlan) FindByPrimaryKey(id string) (result JourneyPlan, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *JourneyPlan) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *JourneyPlan) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *JourneyPlan) FindAll(whereCondition []database.WhereCondition) (results []JourneyPlan, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *JourneyPlan) Save() (result JourneyPlan, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
