package login

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type LoginUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type LoginUserResponse struct {
	gorm.Model
	UUID          uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid"`
	Username      string    `json:"username" form:"username" gorm:"column:username;comment:用户名"`
	Email         string    `json:"email" form:"email" gorm:"column:email;comment:电子邮箱"`
	Telephone     int       `json:"telephone" form:"telephone" gorm:"column:telephone;comment:手机"`
	EffectiveDate time.Time `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:生效日期"`
	ExpiringDate  time.Time `json:"expiringDate" form:"expiringDate" gorm:"column:expiring_date;comment:失效日期"`
	Status        bool      `json:"status" form:"status" gorm:"column:status;comment:状态"`
}

type LoginResponse struct {
	Username string    `json:"username"`
	UUID     uuid.UUID `json:"uuid"`
	// Name     string    `json:"name"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"ExpiresAt"`
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
