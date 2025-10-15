package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	coreModel "infopack.co.in/offybox/app/models/core"
	"time"
)

// DistributorUser [...]
type DistributorUser struct {
	ID            string         `gorm:"primaryKey;column:id" json:"-"`
	DistributorID string         `gorm:"column:distributor_id" json:"distributorId"`
	Distributor   Distributor    `gorm:"joinForeignKey:distributor_id;foreignKey:id;references:DistributorID" json:"entDistributorList"`
	UserID        string         `gorm:"column:user_id" json:"userId"`
	CoreUser      coreModel.User `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	RoleID        string         `gorm:"column:role_id" json:"roleId"`
	CoreRole      coreModel.Role `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID" json:"coreRoleList"`
	Status        string         `gorm:"column:status" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *DistributorUser) TableName() string {
	return "ent_distributor_user"
}

// DistributorUserColumns get sql column name.获取数据库列名
var DistributorUserColumns = struct {
	ID            string
	DistributorID string
	UserID        string
	RoleID        string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	DistributorID: "distributor_id",
	UserID:        "user_id",
	RoleID:        "role_id",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

func (m *DistributorUser) FindByPrimaryKey(id string) (result DistributorUser, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *DistributorUser) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = StatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *DistributorUser) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *DistributorUser) FindAll(whereCondition []database.WhereCondition) (results []DistributorUser, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *DistributorUser) Save() (result DistributorUser, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *DistributorUser) FindByUserId(userId string) (result DistributorUser, err error) {
	err = database.MysqlDB.Model(m).Where("`user_id` = ?", userId).First(&result).Error
	return
}

func (m *DistributorUser) FindOneByCondition(condition []database.WhereCondition) (result DistributorUser, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &condition, nil, nil)
	err = db.First(&result).Error
	return
}
