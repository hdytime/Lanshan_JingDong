package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24

var MySecret = []byte("蓝山寒假考核伪京东")
