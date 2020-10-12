package sysRoleMenus

import (
	"errors"
	"god2admin/global"
	"god2admin/public/request"
	"time"

	"gorm.io/gorm"
)

var dateDisable = time.Date(2099, 12, 31, 0, 0, 0, 0, time.Local)

// 用户角色分配
func roleMenusList(page request.PageListInfo, where map[string]interface{}) (roleMenusList []SysRoleMenus, total int64, err error) {
	limit := page.Size
	offset := page.Size * (page.Current - 1)
	// db := global.DB.Model(&SysRoleMenus{})
	db := global.DB.Table("sys_role_menus_v")
	err = db.Where(where).Count(&total).Error
	err = db.Where(where).Order("status, role_id, menu_id").Limit(limit).Offset(offset).Find(&roleMenusList).Error
	return
}

func roleMenusTreeMap() (treeMap map[*SysRoleMenusAssignKey][]SysRoleMenusAssignList, err error) {
	var menuList []SysRoleMenusAssignList
	treeMap = make(map[*SysRoleMenusAssignKey][]SysRoleMenusAssignList)
	err = global.DB.Table("sys_role_menus_v").Find(&menuList).Error
	for _, v := range menuList {
		treeMap[&SysRoleMenusAssignKey{v.RoleId, v.ParentId}] = append(treeMap[&SysRoleMenusAssignKey{v.RoleId, v.ParentId}], v)
	}
	return
}

func roleMenusAssign(sysRoleMenusAssign SysRoleMenusAssign) (err error) {
	db := global.DB.Model(&SysRoleMenus{})
	for _, v := range sysRoleMenusAssign.MenuId {
		sysRoleMenus := SysRoleMenus{
			RoleId:      sysRoleMenusAssign.RoleId,
			MenuId:      v,
			Status:      true,
			DateEnable:  time.Now(),
			DateDisable: dateDisable,
		}
		err = db.Create(&sysRoleMenus).Error
	}
	return
}

func roleMenusUpdate(sysRoleMenus *SysRoleMenus) (err error) {
	var rm SysRoleMenus
	if errors.Is(global.DB.Model(&SysRoleMenus{}).Where("id = ?", sysRoleMenus.ID).First(&rm).Error, gorm.ErrRecordNotFound) {
		return errors.New("更新失败，无此分配行！")
	}
	// 更新停用日期
	today := time.Now()
	if sysRoleMenus.Status {
		sysRoleMenus.DateDisable = dateDisable
		userRolesEnableParent(sysRoleMenus)
	} else if today.Before(sysRoleMenus.DateDisable) {
		// 判断数据库中失效日期是否大于今天
		sysRoleMenus.DateDisable = today
		userRolesDisableChildren(sysRoleMenus)
	}
	err = global.DB.Save(sysRoleMenus).Error
	return
}

// 失效子菜单分配
func userRolesDisableChildren(sysRoleMenus *SysRoleMenus) (err error) {
	var rmal []SysRoleMenusAssignList
	err = global.DB.Table("sys_role_menus_v").Where("menu_status = 1 AND role_id = ? AND parent_id = ?", sysRoleMenus.RoleId, sysRoleMenus.MenuId).Scan(&rmal).Error
	if err != nil {
		return
	}
	for _, v := range rmal {
		var sysRoleMenus SysRoleMenus
		err = global.DB.Model(&SysRoleMenus{}).Where("id = ?", v.ID).First(&sysRoleMenus).Error
		global.DB.Model(&sysRoleMenus).Updates(map[string]interface{}{"status": 0, "date_disable": time.Now()})
		userRolesDisableChildren(&sysRoleMenus)
	}
	return
}

// 生效上级菜单分配
func userRolesEnableParent(sysRoleMenus *SysRoleMenus) (err error) {
	var rmal SysRoleMenusAssignList
	err = global.DB.Table("sys_role_menus_v").Where("role_id = ? AND menu_id = ?", sysRoleMenus.RoleId, sysRoleMenus.MenuId).Scan(&rmal).Error
	var rm SysRoleMenus
	err = global.DB.Model(&SysRoleMenus{}).Where("status = 0 AND role_id = ? AND menu_id = ?", rmal.RoleId, rmal.ParentId).First(&rm).Error
	if err != nil {
		return
	}
	global.DB.Model(&rm).Updates(map[string]interface{}{"status": 1, "date_disable": dateDisable})
	userRolesEnableParent(&rm)
	return
}
