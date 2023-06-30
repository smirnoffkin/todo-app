package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/smirnoffkin/todo-app/internal/config"
	"github.com/smirnoffkin/todo-app/internal/repository"
	"github.com/smirnoffkin/todo-app/pkg/models"
)

const expireAt = time.Hour

var salt = config.Settings.SaltForHashPassword

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = getPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func getPasswordHash(password string) string {
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashedPassword.Sum([]byte(salt)))
}

func (s *AuthService) CreateAccessToken(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email, getPasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(expireAt),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(config.Settings.SecretKey))
}

func (s *AuthService) VerifyAccessToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.Settings.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
