package customerrors

import "errors"

var (
	ErrPromoCodeNotFound = errors.New("promo code not found")
	ErrPromoCodeExists   = errors.New("promo code already exists")
)
