package sysUserRoles

import (
	"errors"
	"fmt"
	"god2admin/global"
	"god2admin/public/request"
	"time"

	"gorm.io/gorm"
)

// 用户角色分配
func userRolesList(page request.PageListInfo, where map[string]interface{}) (userRolesList []SysUserRoles, total int64, err error) {
	limit := page.Size
	offset := page.Size * (page.Current - 1)
	db := global.DB.Model(&SysUserRoles{})
	err = db.Where(where).Count(&total).Error
	err = db.Where(where).Order("status, user_id, role_id, main_role").Limit(limit).Offset(offset).Find(&userRolesList).Error
	return
}

func userRolesAssign(sysUserRolesAssign SysUserRolesAssign) (err error) {
	fmt.Println(sysUserRolesAssign.RoleId)
	for _, v := range sysUserRolesAssign.RoleId {
		var ur SysUserRoles
		if errors.Is(global.DB.Model(&SysUserRoles{}).Where("user_id = ? AND role_id = ?", sysUserRolesAssign.UserId, v).First(&ur).Error, gorm.ErrRecordNotFound) {
			fmt.Println("create", v)
			sysUserRoles := SysUserRoles{
				UserId:      sysUserRolesAssign.UserId,
				RoleId:      v,
				Status:      true,
				DateEnable:  time.Now(),
				DateDisable: time.Date(2099, 12, 31, 0, 0, 0, 0, time.Local),
				MainRole:    v == 2,
			}
			err = global.DB.Model(&SysUserRoles{}).Create(&sysUserRoles).Error
		}
	}
	return
}

func userRolesUpdate(sysUserRoles *SysUserRoles) (err error) {
	var ur SysUserRoles
	err = global.DB.Model(&SysUserRoles{}).Where("id = ?", sysUserRoles.ID).First(&ur).Error
	if err != nil {
		return errors.New("更新失败，无此分配行！")
	}

	// 如果新记录状态为true，旧记录状态为false,旧记录主要角色为true，
	// 更新新记录主要角色为false。
	//   如果新记录主要角色为true，旧记录主要角色为false，
	//   查询主要角色为true的记录,更新为false
	//   如果新记录主要角色为false，旧记录主要角色为true，
	//   返回异常：主要角色无法变更。请先设置其他角色为主要角色。

	// 如果新记录状态为false，旧记录主要角色为true，
	// 返回异常：主要角色无法停用。请先设置其他角色为主要角色。
	// 如果新记录状态为false，旧记录主要角色为false，
	// 更新新记录主要角色为false。

	// 更新停用日期
	today := time.Now()
	if sysUserRoles.Status {
		sysUserRoles.DateDisable = time.Date(2099, 12, 31, 0, 0, 0, 0, time.Local)
	} else if today.Before(sysUserRoles.DateDisable) {
		// 判断数据库中失效日期是否大于今天
		sysUserRoles.DateDisable = today
	}
	err = global.DB.Save(sysUserRoles).Error
	return
}
