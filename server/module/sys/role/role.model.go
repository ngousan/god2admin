package sysRole

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysRole struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"        gorm:"column:uuid;        comment:uuid"`
	Name        string    `json:"name"        gorm:"column:name;        comment:名称"`
	Title       string    `json:"title"       gorm:"column:title;       comment:标题"`
	Description string    `json:"description" gorm:"column:description; comment:描述"`
	Status      bool      `json:"status"      gorm:"column:status;      comment:状态"`
}

func (SysRole) TableName() string {
	return "sys_roles_b"
}
