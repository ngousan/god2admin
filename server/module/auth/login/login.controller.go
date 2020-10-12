package login

import (
	"fmt"
	"god2admin/global"
	"god2admin/middleware"
	authJwt "god2admin/module/auth/jwt"
	sysUser "god2admin/module/sys/user"
	"god2admin/public/request"
	"god2admin/public/response"
	"god2admin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkDetailed(CaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

func Login(c *gin.Context) {
	var L LoginUser
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	fmt.Println(L.Captcha)
	fmt.Println(L.CaptchaId)
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &LoginUser{Username: L.Username, Password: L.Password}
		if user, err := login(U); err != nil {
			response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			fmt.Println("asdf")
			issueToken(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}

}

// 登录以后签发jwt
func issueToken(c *gin.Context, user sysUser.SysUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID: user.UUID,
		ID:   user.ID,
		// NickName:    user.NickName,
		// AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "God2admin",                    // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	// userMenuList := menus.UserMenusList()
	if !global.CONFIG.System.UseMultipoint {
		response.OkWithData(LoginResponse{
			Username: user.Username,
			UUID:     user.UUID,
			// Name:     user.Username,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			// Menus:     []menus.AuthUserMenus{userMenuList},
		}, c)
		return
	}
	var loginJwt authJwt.JwtBlacklist
	loginJwt.Jwt = token
	err, jwtStr := authJwt.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := authJwt.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(LoginResponse{
			Username: user.Username,
			UUID:     user.UUID,
			// Name:     user.Username,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			// Menus:     []menus.AuthUserMenus{userMenuList},
		}, c)
	} else if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT authJwt.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := authJwt.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := authJwt.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(LoginResponse{
			Username: user.Username,
			UUID:     user.UUID,
			// Name:     user.Username,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			// Menus:     []menus.AuthUserMenus{userMenuList},
		}, c)
	}
}
