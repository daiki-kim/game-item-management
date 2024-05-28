package dtos

type NewItemDTO struct {
	Name        string `json:"name" binding:"required,min=2,max=20"`
	Description string `json:"description"`
}

type UpdateItemDTO struct {
	Name        *string `json:"name" binding:"omitnil"`
	Description *string `json:"description" binding:"omitnil"`
}
