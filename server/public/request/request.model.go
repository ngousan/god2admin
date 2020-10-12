package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type PageListInfo struct {
	Size    int `json:"size"`
	Current int `json:"current"`
}

type CustomClaims struct {
	UUID uuid.UUID
	ID   uint
	jwt.StandardClaims
}
