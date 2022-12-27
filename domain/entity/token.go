package entity

import "github.com/golang-jwt/jwt"

type MyCustomClaims struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}
