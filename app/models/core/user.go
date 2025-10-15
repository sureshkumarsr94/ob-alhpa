package core

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/database"
	"time"
)

// User [...]
type User struct {
	ID                 string    `gorm:"primaryKey;column:id" json:"-"`
	FirstName          string    `gorm:"column:first_name" json:"firstName"`
	LastName           string    `gorm:"column:last_name" json:"lastName"`
	Username           string    `gorm:"column:username" json:"username"`
	Type               string    `gorm:"column:type" json:"type"`
	Email              string    `gorm:"column:email" json:"email"`
	Mobile             string    `gorm:"column:mobile" json:"mobile"`
	Password           string    `gorm:"column:password" json:"password"`
	Code               string    `gorm:"column:code" json:"code"`
	AssociateUserId    string    `gorm:"column:associate_user_id" json:"associateUserId"`
	LastPasswordChange null.Time `gorm:"column:last_password_change" json:"lastPasswordChange"`
	Status             string    `gorm:"column:status" json:"status"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "core_user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Type      string
	Email     string
	Mobile    string
	Password  string
	Code      string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	FirstName: "first_name",
	LastName:  "last_name",
	Username:  "username",
	Type:      "type",
	Email:     "email",
	Mobile:    "mobile",
	Password:  "password",
	Code:      "code",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = UserStatusActive
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *User) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *User) Save() (result User, err error) {
	err = database.MysqlDB.Save(&m).Last(&result).Error
	return
}
func (m *User) FindByPrimaryKey(userID string) (result User, err error) {
	err = database.MysqlDB.Model(m).Where("id = ?", userID).Find(&result).Error
	return
}

func (m *User) FindOneByCondition(whereCondition *[]database.WhereCondition, orCondition *[]database.WhereCondition) (result User, err error) {
	db := database.MysqlDB.Model(m)
	db = database.ConditionBuilder(db, whereCondition, orCondition, nil)
	err = db.Find(&result).Error
	return
}

func (m *User) FindByEmailOrMobile(userType string, email string, mobile string) (result User, err error) {
	err = database.MysqlDB.Model(m).
		Where("type = ?", userType).
		Where("( email = ? or mobile =?)", email, mobile).
		Find(&result).Error
	return
}
