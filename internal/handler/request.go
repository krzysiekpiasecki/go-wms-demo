package handler

type CreateOrderRequest struct {
	ProductID int64   `json:"productId" binding:"required,min=1"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Comment   *string `json:"comment"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
