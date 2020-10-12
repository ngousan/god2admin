package sysUser

import (
	"errors"
	"god2admin/global"
	"god2admin/public/request"
	"god2admin/public/response"
	"god2admin/utils"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户管理
func userList(page request.PageListInfo, sysUser SysUser) (userList []SysUser, total int64, err error) {
	limit := page.Size
	offset := page.Size * (page.Current - 1)
	db := global.DB.Model(&SysUser{})
	if sysUser.Username == "" {
		err = db.Count(&total).Error
		// err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
		err = db.Limit(limit).Offset(offset).Find(&userList).Error
	} else {
		err = db.Where("username like ?", sysUser.Username).Count(&total).Error
		// err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
		err = db.Where("username like ?", sysUser.Username).Limit(limit).Offset(offset).Find(&userList).Error
	}
	return
}

func userRegister(sysUser SysUser) (err error) {
	var u SysUser
	if !errors.Is(global.DB.Model(&SysUser{}).Where("username = ?", sysUser.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已被注册！")
	}
	if !errors.Is(global.DB.Model(&SysUser{}).Where("email = ?", sysUser.Email).First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("电子邮箱已被注册！")
	}
	if !errors.Is(global.DB.Model(&SysUser{}).Where("telephone = ?", sysUser.Telephone).First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("移动电话已被注册！")
	}
	sysUser.UUID = uuid.NewV4()
	sysUser.Password = utils.MD5V([]byte(sysUser.Password))
	sysUser.Status = true
	sysUser.DateEnable = time.Now()
	sysUser.DateDisable = time.Date(2099, 12, 31, 0, 0, 0, 0, time.Local)
	err = global.DB.Create(&sysUser).Error
	return
}

func userUpdate(sysUser *SysUser) (err error) {
	var u SysUser
	db := global.DB.Model(&SysUser{})
	err = db.Where("id like ?", sysUser.ID).First(&u).Error
	if err != nil {
		return errors.New("更新失败，无此用户！")
	}
	// 更新停用日期
	today := time.Now()
	if sysUser.Status {
		sysUser.DateDisable = time.Date(2099, 12, 31, 0, 0, 0, 0, time.Local)
	} else if today.Before(sysUser.DateDisable) {
		// 判断数据库中失效日期是否大于今天
		sysUser.DateDisable = today
	}
	err = global.DB.Save(sysUser).Error
	return
}

func userDict() (dict []response.Dict, err error) {
	var u []SysUser
	db := global.DB.Model(&SysUser{})
	err = db.Find(&u).Error
	if err != nil {
		return dict, errors.New("查询用户失败！")
	}
	for _, v := range u {
		dict = append(dict, response.Dict{
			Value: strconv.Itoa(int(v.ID)),
			Label: v.Username,
		})
	}
	return
}
