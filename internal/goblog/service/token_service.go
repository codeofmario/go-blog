package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/model"
	"goblog.com/goblog/internal/goblog/repository"
	"strings"
	"time"
)

type TokenService interface {
	GetAccessToken(userId string, tokenId string) (string, error)
	GetRefreshToken(userId string, tokenId string) (string, error)
	SaveAccessToken(userId string, tokenId string, token string) (bool, error)
	SaveRefreshToken(userId string, tokenId string, token string) (bool, error)
	DeleteAccessTokens(userId string, tokenId string) (bool, error)
	DeleteRefreshTokens(userId string, tokenId string) (bool, error)
	GenerateAccessToken(userId string, tokenId string) (string, error)
	GenerateRefreshToken(userId string, tokenId string) (string, error)

	ExtractFromAuthHeader(authHeader string) (string, error)
	Parse(tokenString string, secret string) (*model.TokenClaims, error)
}

type TokenServiceImpl struct {
	Settings *config.Settings
	Repository repository.TokenRepository
}

func NewTokenService(settings *config.Settings, repository repository.TokenRepository) TokenService {
	return &TokenServiceImpl{Settings: settings, Repository: repository}
}

func (s *TokenServiceImpl) GetAccessToken(userId string, tokenId string) (string, error) {
	key := fmt.Sprintf("%s.%s.at", userId, tokenId)
	return s.Repository.GetToken(key)
}

func (s *TokenServiceImpl) GetRefreshToken(userId string, tokenId string) (string, error) {
	key := fmt.Sprintf("%s.%s.rt", userId, tokenId)
	return s.Repository.GetToken(key)
}

func (s *TokenServiceImpl) SaveAccessToken(userId string, tokenId string, token string) (bool, error) {
	key := fmt.Sprintf("%s.%s.at", userId, tokenId)
	exp := time.Now().Add(time.Minute * 60)
	return s.Repository.SaveToken(key, token, exp)
}

func (s *TokenServiceImpl) SaveRefreshToken(userId string, tokenId string, token string) (bool, error) {
	key := fmt.Sprintf("%s.%s.rt", userId, tokenId)
	exp := time.Now().Add(time.Hour * 24 * 7)
	return s.Repository.SaveToken(key, token, exp)
}

func (s *TokenServiceImpl) DeleteAccessTokens(userId string, tokenId string) (bool, error) {
	key := fmt.Sprintf("%s.%s.at", userId, tokenId)
	return s.Repository.DeleteToken(key)
}

func (s *TokenServiceImpl) DeleteRefreshTokens(userId string, tokenId string) (bool, error) {
	key := fmt.Sprintf("%s.%s.rt", userId, tokenId)
	return s.Repository.DeleteToken(key)
}

func (s *TokenServiceImpl) GenerateAccessToken(userId string, tokenId string) (string, error) {
	claims := model.TokenClaims{}
	claims.Issuer = "goblog"
	claims.Subject = userId
	claims.TokenId = tokenId
	claims.IsRefresh = false
	claims.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Minute * 60)}
	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Settings.AccessSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenServiceImpl) GenerateRefreshToken(userId string, tokenId string) (string, error) {
	claims := model.TokenClaims{}
	claims.Issuer = "goblog"
	claims.Subject = userId
	claims.TokenId = tokenId
	claims.IsRefresh = true
	claims.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24 * 7)}
	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Settings.RefreshSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenServiceImpl) ExtractFromAuthHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.InternalServerError{Msg: "Empty header."}
	}

	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.InternalServerError{Msg: "Incorrect format."}
	}

	return parts[1], nil
}

func (s *TokenServiceImpl) Parse(tokenString string, secret string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.BadRequestError{Msg: fmt.Sprintf("Signing method %v is incorrect", token.Header["alg"])}
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.InternalServerError{Msg: "Invalid token"}
}
