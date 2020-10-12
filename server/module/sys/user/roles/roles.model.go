package sysUserRoles

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 角色分配
type SysUserRoles struct {
	gorm.Model
	UserId      int       `json:"userId"      gorm:"column:user_id;     comment:用户ID"`
	RoleId      int       `json:"roleId"      gorm:"column:role_id;     comment:角色ID"`
	Status      bool      `json:"status"      gorm:"column:status;       comment:状态"`
	MainRole    bool      `json:"mainRole"    gorm:"column:main_role;    comment:主要角色"`
	DateEnable  time.Time `json:"dateEnable"  gorm:"column:date_enable;  comment:启用日期"`
	DateDisable time.Time `json:"dateDisable" gorm:"column:date_disable; comment:停用日期"`
	// User        SysUser         `json:"user"        gorm:"ForeignKey:User_id;  comment:用户"`
	// Role        sysRole.SysRole `json:"role"        gorm:"ForeignKey:Role_id;  comment:角色"`
}

func (SysUserRoles) TableName() string {
	return "sys_user_roles_r"
}

type SysUserRolesAssign struct {
	UserId int   `json:"userId"`
	RoleId []int `json:"roleId"`
}
