package requests


type TransferRequest struct {
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}