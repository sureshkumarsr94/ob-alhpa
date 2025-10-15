package core

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/database"
	"time"
)

// UserCredentialRequest [...]
type UserCredentialRequest struct {
	ID        string    `gorm:"primaryKey;column:id" json:"-"`
	UserID    string    `gorm:"column:user_id" json:"userId"`
	User      User      `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"coreUserList"`
	Mode      string    `gorm:"column:mode" json:"mode"`
	ExpireAt  time.Time `gorm:"column:expire_at" json:"expireAt"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name.
func (m *UserCredentialRequest) TableName() string {
	return "core_user_credential_request"
}

// UserCredentialRequestColumns get sql column name.
var UserCredentialRequestColumns = struct {
	ID        string
	UserID    string
	Mode      string
	ExpireAt  string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Mode:      "mode",
	ExpireAt:  "expire_at",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (m *UserCredentialRequest) BeforeCreate(tx *gorm.DB) (err error) {

	m.ID = uuid.New().String()
	m.Status = CredentialRequestStatusInitiated
	m.CreatedAt = time.Now()
	m.ExpireAt = time.Now().Add(24 * time.Hour)

	return
}

func (m *UserCredentialRequest) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *UserCredentialRequest) Save() (result UserCredentialRequest, err error) {
	err = database.MysqlDB.Create(&m).Preload("User").Last(&result).Error
	return
}

func (m *UserCredentialRequest) GetResetLink() string {
	var platform string

	if m.User.Type == constants.UserTypeDistributor {
		platform = "cfg.TenantConfig.PartnerPortalUrl"
	} else if m.User.Type == constants.UserTypeEmployee {
		platform = "cfg.TenantConfig.EmployeePortalUrl"
	}

	return fmt.Sprintf("%s/reset-password/%s", platform, m.ID)
}

func (m *UserCredentialRequest) FindByPrimaryKey(requestId string) (result UserCredentialRequest, err error) {
	err = database.MysqlDB.Model(m).Where("`id` = ?", requestId).Preload("User").Find(&result).Error
	return
}
