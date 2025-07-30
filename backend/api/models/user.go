package models

type User struct { // Структура пользователя
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // e.g., "admin", "user"

}

type SignUpInput struct { // Структура для регистрации пользователя
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct { // Структура для входа пользователя
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserReset struct { // Структура для сброса пароля
	Token   string `json:"token" binding:"required"`
	NewPass string `json:"new_password" binding:"required"`
}

type ResetRequest struct { // Структура для запроса сброса пароля
	Login string `json:"login" binding:"required"`
}

type UserInfo struct { // Структура для получения информации о пользователе
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
