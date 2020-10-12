package sysMenu

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysMenu struct {
	gorm.Model
	UUID         uuid.UUID `json:"uuid"         gorm:"column:uuid;          comment:uuid"`
	Name         string    `json:"name"         gorm:"column:name;          comment:菜单"`
	Title        string    `json:"title"        gorm:"column:title;         comment:标题"`
	Description  string    `json:"description"  gorm:"column:description;   comment:描述"`
	Path         string    `json:"path"         gorm:"column:path;          comment:路径"`
	Icon         string    `json:"icon"         gorm:"column:icon;          comment:图标"`
	ParentId     int       `json:"parentId"     gorm:"column:parent_id;     comment:上级"`
	RedirectName string    `json:"redirectName" gorm:"column:redirect_name; comment:重定向"`
	Component    string    `json:"component"    gorm:"column:component;     comment:组件"`
	Auth         bool      `json:"auth"         gorm:"column:auth;          comment:权限"`
	Cache        bool      `json:"cache"        gorm:"column:cache;         comment:缓存"`
	Status       bool      `json:"status"       gorm:"column:status;        comment:状态"`
	Sort         int       `json:"sort"         gorm:"column:sort;          comment:排序"`
	Children     []SysMenu `json:"children"     gorm:"-"`
}

func (SysMenu) TableName() string {
	return "sys_menus_b"
}

type SysMenuList struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ParentId    int       `json:"parentId"`
	Router      string    `json:"router"`
	Path        string    `json:"path"`
	Component   string    `json:"component"`
	Status      bool      `json:"status"`
	Sort        int       `json:"sort"`
	Icon        string    `json:"icon"`
	Default     bool      `json:"default"`
	KeepAlive   bool      `json:"keepAlive"`
	HasChildren bool      `json:"hasChildren"`
	Children    []SysMenu `json:"children"`
}
