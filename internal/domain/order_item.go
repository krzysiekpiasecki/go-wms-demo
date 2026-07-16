package domain

type OrderItem struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"orderId"`
	ProductID int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}
