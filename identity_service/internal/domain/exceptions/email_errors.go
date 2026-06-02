package exceptions

import "errors"

var (
	ErrInvalidEmail = errors.New("email inválido")
)