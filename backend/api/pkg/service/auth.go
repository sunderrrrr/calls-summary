package service

import (
	"api/models"
	"api/pkg/repository"
	"api/pkg/utils/logger"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

/*
Реализация сервиса аутентификации: отправка запросов в репозиторий, генерация и парсинг токенов
*/

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	salt      = os.Getenv("PASS_SALT")
	secretKey = os.Getenv("JWT_KEY")
)

func saltPassword(password, salt string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(salt + password)) // соль добавляется к паролю
	return hex.EncodeToString(h.Sum(nil))
}

func createHash(password string) (string, error) {
	hash := saltPassword(password, salt)
	return salt + ":" + hash, nil
}

func (s *AuthService) SignUp(input models.SignUpInput) (int, error) {
	passwordHash, _ := createHash(input.Password)
	return s.repo.SignUp(input.Email, input.Name, passwordHash)
}

func (s *AuthService) GenerateToken(input models.SignInInput) (string, error) {
	h, _ := createHash(input.Password)
	user, err := s.repo.GetUser(input.Email, h)
	if err != nil {
		return "", err

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
	return token.SignedString([]byte(secretKey))

}

func (s *AuthService) ParseToken(accessToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		logger.Log.Errorf("AuthService.ParseToken error: %v", err.Error())

		return models.User{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.User{}, errors.New("token claims are not of type *tokenClaims")
	}

	returnUser := models.User{
		Id:    claims.Id,
		Name:  claims.Name,
		Email: claims.Email,
	}

	return returnUser, nil

}
