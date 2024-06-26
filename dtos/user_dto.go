package dtos

type SignupUserDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	Description string `json:"description"`
}

type LoginUserDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	Description string `json:"description"`
}

type GetUsersDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserDTO struct {
	Name        *string `json:"name" binding:"omitnil"`
	Email       *string `json:"email" binding:"omitnil,email"`
	Description *string `json:"description" binding:"omitnil"`
}
