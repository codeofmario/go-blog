package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/dto/response"
	customErrors "goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/model"
	"goblog.com/goblog/internal/goblog/util"
)

type AuthService interface {
	Login(user *model.User) (*response.TokensResponseDto, error)
	Logout(authHeader string) error
	Refresh(tokenString string) (*response.TokensResponseDto, error)
}

type AuthServiceImpl struct {
	UserService  UserService
	TokenService TokenService
	Settings     *config.Settings
}

func NewAuthService(userService UserService, tokenService TokenService, settings *config.Settings) AuthService {
	return &AuthServiceImpl{
		UserService:  userService,
		TokenService: tokenService,
		Settings:     settings,
	}
}

func (s *AuthServiceImpl) Login(user *model.User) (*response.TokensResponseDto, error) {
	dbUser, err := s.UserService.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = s.checkUser(user, dbUser)
	if err != nil {
		return nil, err
	}

	return s.getTokens(dbUser.ID.String())
}

func (s *AuthServiceImpl) Logout(authHeader string) error {
	tokenString, err := s.TokenService.ExtractFromAuthHeader(authHeader)
	if err != nil {
		return err
	}

	claims, err := s.TokenService.Parse(tokenString, s.Settings.AccessSecret)
	if err != nil {
		return err
	} else if claims.IsRefresh {
		return customErrors.BadRequestError{Msg: "Bad Request"}
	}

	err = s.deleteTokens(claims.Subject, claims.TokenId)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) Refresh(tokenString string) (*response.TokensResponseDto, error) {

	claims, err := s.TokenService.Parse(tokenString, s.Settings.RefreshSecret)
	if err != nil {
		return nil, err
	} else if !claims.IsRefresh {
		return nil, customErrors.BadRequestError{Msg: "Bad Request"}
	}

	userId, _ := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	_, err = s.UserService.GetOne(userId)
	if err != nil {
		return nil, err
	}

	err = s.deleteTokens(claims.Subject, claims.TokenId)
	if err != nil {
		return nil, err
	}

	return s.getTokens(claims.Subject)
}

func (s *AuthServiceImpl) checkUser(user *model.User, dbUser *model.User) error {
	email := user.Email
	password, _ := util.HashPassword(user.Password)

	if dbUser.Email != email || !util.CheckPasswordHash(password, dbUser.Password) {
		return errors.New("")
	}

	return nil
}

func (s *AuthServiceImpl) getTokens(userId string) (*response.TokensResponseDto, error) {
	tokenId := uuid.New().String()
	tokens, err := s.generateTokens(userId, tokenId)
	if err != nil {
		return nil, err
	}

	err = s.saveTokens(userId, tokenId, tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *AuthServiceImpl) generateTokens(userId string, tokenId string) (*response.TokensResponseDto, error) {
	accessToken, err := s.TokenService.GenerateAccessToken(userId, tokenId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.TokenService.GenerateRefreshToken(userId, tokenId)
	if err != nil {
		return nil, err
	}

	return &response.TokensResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) saveTokens(userId string, tokenId string, tokens *response.TokensResponseDto) error {
	_, err := s.TokenService.SaveAccessToken(userId, tokenId, tokens.AccessToken)
	if err != nil {
		return err
	}

	_, err = s.TokenService.SaveRefreshToken(userId, tokenId, tokens.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) deleteTokens(userId string, tokenId string) error {
	_, err := s.TokenService.DeleteAccessTokens(userId, tokenId)
	if err != nil {
		return err
	}

	_, err = s.TokenService.DeleteRefreshTokens(userId, tokenId)
	if err != nil {
		return err
	}

	return nil
}
