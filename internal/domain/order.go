package domain

import "time"

type Order struct {
	ID        int64
	Status    string
	Comment   *string
	CreatedAt time.Time
	Items     []OrderItem
}
