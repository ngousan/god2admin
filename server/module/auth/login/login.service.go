package login

import (
	"god2admin/global"
	sysUser "god2admin/module/sys/user"
	"god2admin/utils"
)

func login(loginUser *LoginUser) (reUser *sysUser.SysUser, err error) {
	var user sysUser.SysUser
	loginUser.Password = utils.MD5V([]byte(loginUser.Password))
	// err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	err = global.DB.Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&user).Error
	user.Password = ""
	return &user, err
}
