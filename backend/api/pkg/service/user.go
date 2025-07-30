package service

import (
	"api/models"
	"api/pkg/repository"
	"api/pkg/utils/logger"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

/*
Операции над пользовательскими данными. Пока что реализовано изменение пароля
*/

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

type ResetClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func generateResetToken(email string) (string, error) {
	if email == "" {
		return "", errors.New("email is required")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ResetClaims{

		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(resetTokenTTL).Unix(),
		},
		Email: email,
	})
	return token.SignedString([]byte(secretKey))
}

func (s *UserService) ForgotPassword(request models.ResetRequest) error {
	token, err := generateResetToken(request.Login)
	if err != nil {
		return err
	}
	resetLink := fmt.Sprintf("%s/forgot?t=%s", os.Getenv("FRONTEND_URL"), token)
	logger.Log.Println(resetLink)
	// Далее будет реализация отправки письма
	return nil
}

func (s *UserService) ResetPassword(request models.UserReset) error {
	token, err := jwt.ParseWithClaims(request.Token, &ResetClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*ResetClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}
	newHash, err := createHash(request.NewPass)
	return s.repo.ResetPassword(claims.Email, newHash)
}
