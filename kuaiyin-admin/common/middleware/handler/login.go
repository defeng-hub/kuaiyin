package handler

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"
)

type LoginType int
type LoginTypeDto struct {
	LoginType LoginType `form:"loginType" json:"loginType"`
	Overload  bool      `form:"overload" json:"overload"` //是否重载，是否慢速登录
}

var (
	Normal LoginType = 0
	SSO    LoginType = 1
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
	LoginTypeDto
}
type LoginPcNumberDto struct {
	PcNumber string `form:"pcNumber" json:"pcNumber" binding:"required"`
	LoginTypeDto
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = '2'", u.Username).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		return
	}
	err = tx.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		log.Errorf("get role error, %s", err.Error())
		return
	}
	return
}

// 通过警官号查人，再登录
func (u *LoginPcNumberDto) GetUserByPcNumber(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("pc_number = ?  and status = '2'", u.PcNumber).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	err = tx.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		log.Errorf("get role error, %s", err.Error())
		return
	}
	return
}
