package sysRole

import (
	"errors"
	"god2admin/global"
	"god2admin/public/request"
	"god2admin/public/response"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func list(page request.PageListInfo) (roleList []SysRole, total int64, err error) {
	limit := page.Size
	offset := page.Size * (page.Current - 1)
	db := global.DB.Model(&SysRole{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&roleList).Error
	return
}

func insert(sysRole SysRole) (err error) {
	db := global.DB.Model(&SysRole{})
	sysRole.UUID = uuid.NewV4()
	sysRole.Status = true
	err = db.Create(&sysRole).Error
	return err
}

func update(sysRole *SysRole) (err error) {
	db := global.DB.Model(&SysRole{})
	err = db.Save(sysRole).Error
	return err
}

func roleDict() (dict []response.Dict, err error) {
	var r []SysRole
	db := global.DB.Model(&SysRole{})
	err = db.Find(&r).Error
	if err != nil {
		return dict, errors.New("查询角色失败！")
	}
	for _, v := range r {
		dict = append(dict, response.Dict{
			Value: strconv.Itoa(int(v.ID)),
			Label: v.Name,
		})
	}
	return
}
