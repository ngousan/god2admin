package authJwt

import (
	"errors"
	"god2admin/global"

	"gorm.io/gorm"
)

func IsBlacklist(jwt string) bool {
	var jwtBlacklist JwtBlacklist
	// isNotFound := global.DB.Where("jwt = ?", jwt).First().RecordNotFound()
	if errors.Is(global.DB.Where("jwt = ?", jwt).First(&jwtBlacklist).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func JsonInBlacklist(jwtBlacklist JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtBlacklist).Error
	return
}

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(userName).Result()
	return err, redisJWT
}

func SetRedisJWT(jwtList JwtBlacklist, userName string) (err error) {
	err = global.REDIS.Set(userName, jwtList.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
