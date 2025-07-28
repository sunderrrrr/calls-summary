package service_test

import (
	"api/models"
	"api/pkg/service"
	"testing"
)

// fakeAuthRepo реализует интерфейс repository.Auth для тестов.
type fakeAuthRepo struct{}

// Реализация регистрации. Возвращает идентификатор пользователя.
func (r *fakeAuthRepo) SignUp(email, name, password string) (int, error) {
	return 1, nil
}

// Реализация получения пользователя. Возвращает тестового пользователя.
func (r *fakeAuthRepo) GetUser(email, password string) (models.User, error) {
	return models.User{
		Id:    1,
		Name:  "Test User",
		Email: email,
	}, nil
}

func TestSignUp(t *testing.T) {
	repo := &fakeAuthRepo{}
	authService := service.NewAuthService(repo)

	input := models.SignUpInput{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password",
	}

	id, err := authService.SignUp(input)
	if err != nil {
		t.Fatalf("Ошибка при регистрации: %v", err)
	}
	if id != 1 {
		t.Fatalf("Ожидался id 1, получен %d", id)
	}
}

func TestGenerateAndParseToken(t *testing.T) {
	repo := &fakeAuthRepo{}
	authService := service.NewAuthService(repo)

	input := models.SignInInput{
		Email:    "test@example.com",
		Password: "password",
	}

	// Генерация токена.
	token, err := authService.GenerateToken(input)
	if err != nil {
		t.Fatalf("GenerateToken вернула ошибку: %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken вернула пустой токен")
	}

	// Разбор токена.
	user, err := authService.ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken вернула ошибку: %v", err)
	}

	// Проверка соответствия данных.
	if user.Email != input.Email {
		t.Fatalf("Ожидался email %s, получен %s", input.Email, user.Email)
	}
	if user.Name != "Test User" {
		t.Fatalf("Ожидалось имя 'Test User', получено %s", user.Name)
	}
}
