package dtos

type UpdateTradeDTO struct {
	Is_Accepted bool `json:"is_accepted" binding:"required"`
}
