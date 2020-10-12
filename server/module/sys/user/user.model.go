package sysUser

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"        gorm:"column:uuid;         comment:uuid"`
	Username    string    `json:"username"    gorm:"column:username;     unique_index; comment:用户名"`
	Password    string    `json:"password"    gorm:"column:password;     comment:密码"`
	Email       string    `json:"email"       gorm:"column:email;        unique_index;comment:电子邮箱"`
	Telephone   string    `json:"telephone"   gorm:"column:telephone;    unique_index;comment:手机"`
	DateEnable  time.Time `json:"dateEnable"  gorm:"column:date_enable;  comment:启用日期"`
	DateDisable time.Time `json:"dateDisable" gorm:"column:date_disable; comment:停用日期"`
	Status      bool      `json:"status"      gorm:"column:status;       comment:状态"`
	// UserRoles   []SysUserRoles `json:"userRoles"   gorm:"foreignKey:UserId"`
}

func (SysUser) TableName() string {
	return "sys_users_b"
}
