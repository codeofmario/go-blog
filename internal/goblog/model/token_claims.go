package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	TokenId string `json:"tokenId"`
	IsRefresh bool `json:"isRefresh"`
	jwt.RegisteredClaims
}
