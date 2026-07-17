package domain

import "time"

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}
