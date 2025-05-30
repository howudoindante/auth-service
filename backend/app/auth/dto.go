package auth

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
}

type AuthUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
