package domain

import "errors"

var ErrProductNotFound = errors.New("product not found")
var ErrInvalidQuantity = errors.New("invalid quantity")
var ErrInsufficientStock = errors.New("insufficient stock")
