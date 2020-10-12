package sysMenu

import (
	"errors"
	"fmt"
	"god2admin/global"
	"god2admin/public/request"
	"god2admin/public/response"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 获取分页列表，非树形结构
func list(page request.PageListInfo) (menuList []SysMenu, total int64, err error) {
	limit := page.Size
	offset := page.Size * (page.Current - 1)
	db := global.DB.Model(&SysMenu{})
	err = db.Count(&total).Error
	// err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	err = db.Limit(limit).Offset(offset).Find(&menuList).Error
	return
}

// 获取树形结构，返回map
func listTreeMap() (treeMap map[int][]SysMenu, err error) {
	db := global.DB.Model(&SysMenu{})
	var menuList []SysMenu
	treeMap = make(map[int][]SysMenu)
	err = db.Order("sort").Find(&menuList).Error
	for _, v := range menuList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return
}

func insert(sysMenu SysMenu) (err error) {
	var m SysMenu
	fmt.Println(sysMenu)
	if !errors.Is(global.DB.Model(&SysMenu{}).Where("name = ?", sysMenu.Name).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单名重复！")
	}
	if !errors.Is(global.DB.Model(&SysMenu{}).Where("title = ? and parent_id = ?", sysMenu.Title, sysMenu.ParentId).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同菜单标题在同级菜单中！")
	}
	sysMenu.UUID = uuid.NewV4()
	sysMenu.Status = true
	sysMenu.Cache = true
	err = global.DB.Create(&sysMenu).Error
	return
}

func update(sysMenu *SysMenu) (err error) {
	var originalSysMenu SysMenu
	updateMap := make(map[string]interface{})
	updateMap["title"] = sysMenu.Title
	updateMap["description"] = sysMenu.Description
	updateMap["parent_id"] = sysMenu.ParentId
	updateMap["path"] = sysMenu.Path
	updateMap["component"] = sysMenu.Component
	updateMap["icon"] = sysMenu.Icon
	updateMap["sort"] = sysMenu.Title
	updateMap["status"] = sysMenu.Status
	db := global.DB.Where("id = ?", sysMenu.ID).Find(&originalSysMenu)
	err = db.Updates(updateMap).Error
	return
}

func dictTreeMap() (treeMap map[string][]response.DictTree, err error) {
	db := global.DB.Model(&SysMenu{})
	var menuList []SysMenu
	treeMap = make(map[string][]response.DictTree)
	err = db.Order("sort").Find(&menuList).Error
	for _, v := range menuList {
		treeMap[strconv.Itoa(v.ParentId)] = append(treeMap[strconv.Itoa(v.ParentId)], response.DictTree{
			Value: strconv.Itoa(int(v.ID)),
			Label: v.Title,
		})
	}
	return
}
