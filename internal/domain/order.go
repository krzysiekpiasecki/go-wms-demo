package domain

import "time"

type Order struct {
	ID        int64       `json:"id"`
	Status    string      `json:"status"`
	Comment   *string     `json:"comment"`
	CreatedAt time.Time   `json:"createdAt"`
	Items     []OrderItem `json:"items"`
}
