package service

import (
	"context"
	"errors"
	"eticket-test/internal/pkg/jwt"
	"eticket-test/modules/auth/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.AuthUserRepository
	jwt      jwt.JWT
}

func NewAuthService(userRepo repository.AuthUserRepository, jwt jwt.JWT) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

type TokenResponse struct {
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int64       `json:"expires_in"`
	User        interface{} `json:"user"`
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*TokenResponse, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("invalid password")
	}

	// buat claims
	claims := map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
	}

	// generate token pakai pkg/jwt
	signedToken, err := s.jwt.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken: signedToken,
		TokenType:   "Bearer",
		ExpiresIn:   int64(7 * 24 * 3600), // sesuai pkg/jwt kamu (mingguan)
		User: map[string]interface{}{
			"user_id":  user.ID,
			"username": user.Username,
			"fullname": user.FullName,
			"role":     user.Role,
		},
	}, nil
}
