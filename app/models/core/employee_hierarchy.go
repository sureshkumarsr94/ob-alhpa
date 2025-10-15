package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// EmployeeHierarchy [...]
type EmployeeHierarchy struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`
	EmployeeID   string    `gorm:"column:employee_id" json:"employeeId"`
	Employee     Employee  `gorm:"joinForeignKey:employee_id;foreignKey:id;references:EmployeeID" json:"coreEmployeeList"`
	SupervisorID string    `gorm:"column:supervisor_id" json:"supervisorId"`
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *EmployeeHierarchy) TableName() string {
	return "core_employee_hierarchy"
}

// EmployeeHierarchyColumns get sql column name.获取数据库列名
var EmployeeHierarchyColumns = struct {
	ID           string
	EmployeeID   string
	SupervisorID string
	Status       string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	EmployeeID:   "employee_id",
	SupervisorID: "supervisor_id",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

func (m *EmployeeHierarchy) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *EmployeeHierarchy) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *EmployeeHierarchy) Save() (result EmployeeHierarchy, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}

func (m *EmployeeHierarchy) FindAll(whereCondition []database.WhereCondition) (results []EmployeeHierarchy, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, &whereCondition, nil, nil)
	err = db.Find(&results).Error
	return
}

func (m *EmployeeHierarchy) FindByPrimaryKey(id string) (result EmployeeHierarchy, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", id).Find(&result).Error
	return
}

func (m *EmployeeHierarchy) FindByEmployeeSupervisor(employeeID string, supervisorID string) (result EmployeeHierarchy, err error) {
	err = database.MysqlDB.Model(m).
		Where("`employee_id` = ?", employeeID).
		Where("`supervisor_id` = ?", supervisorID).
		Find(&result).Error
	return
}
