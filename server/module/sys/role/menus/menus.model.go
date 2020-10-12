package sysRoleMenus

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 菜单分配
type SysRoleMenus struct {
	gorm.Model
	RoleId      int       `json:"roleId"      gorm:"column:role_id;      comment:角色ID"`
	MenuId      int       `json:"menuId"      gorm:"column:menu_id;      comment:菜单ID"`
	Status      bool      `json:"status"      gorm:"column:status;       comment:状态"`
	DateEnable  time.Time `json:"dateEnable"  gorm:"column:date_enable;  comment:启用日期"`
	DateDisable time.Time `json:"dateDisable" gorm:"column:date_disable; comment:停用日期"`
}

func (SysRoleMenus) TableName() string {
	return "sys_role_menus_r"
}

type SysRoleMenusAssign struct {
	RoleId int   `json:"roleId"`
	MenuId []int `json:"menuId"`
}

type SysRoleMenusAssignKey struct {
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}

type SysRoleMenusAssignList struct {
	ID          uint                     `json:"ID"`
	RoleId      int                      `json:"roleId"`
	MenuId      int                      `json:"menuId"`
	Status      bool                     `json:"status"`
	DateEnable  time.Time                `json:"dateEnable"`
	DateDisable time.Time                `json:"dateDisable"`
	ParentId    int                      `json:"parentId"`
	Children    []SysRoleMenusAssignList `json:"children" gorm:"-"`
}
