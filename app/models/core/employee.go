package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// Employee [...]
type Employee struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	User      User      `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"user"`
	Code      string    `gorm:"column:code" json:"code"`
	Name      string    `gorm:"column:name" json:"name"`
	Mobile    string    `gorm:"column:mobile" json:"mobile"`
	Email     string    `gorm:"column:email" json:"email"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Employee) TableName() string {
	return "core_employee"
}

// EmployeeColumns get sql column name.获取数据库列名
var EmployeeColumns = struct {
	ID        string
	UserID    string
	Code      string
	Name      string
	Mobile    string
	Email     string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Code:      "code",
	Name:      "name",
	Mobile:    "mobile",
	Email:     "email",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (m *Employee) FindByPrimaryKey(id string) (result Employee, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Employee) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Employee) FindAll(whereCondition []database.WhereCondition) (results []Employee, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *Employee) Save() (result Employee, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *Employee) FindByCode(code string) (result Employee, err error) {
	err = database.MysqlDB.Model(m).Where("`code` = ?", code).Find(&result).Error
	return
}
