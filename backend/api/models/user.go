package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // e.g., "admin", "user"

}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReset struct {
	Token   string `json:"token" binding:"required"`
	NewPass string `json:"new_password" binding:"required"`
}

type ResetRequest struct {
	Login string `json:"login" binding:"required"`
}

type UserInfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
