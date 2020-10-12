package authJwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	UUID uuid.UUID
	ID   uint
	jwt.StandardClaims
}

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:'jwt'"`
}
