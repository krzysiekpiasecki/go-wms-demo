package handler

type CreateOrderRequest struct {
	Comment *string                  `json:"comment"`
	Items   []CreateOrderItemRequest `json:"items" binding:"required"`
}

type CreateOrderItemRequest struct {
	ProductID int64 `json:"productId" binding:"required,min=1"`
	Quantity  int   `json:"quantity" binding:"required,min=1"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type CreateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,min=0"`
}
