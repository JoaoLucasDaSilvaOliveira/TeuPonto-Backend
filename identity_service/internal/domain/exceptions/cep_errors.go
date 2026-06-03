package exceptions

import "errors"

var (
	ErrInvalidCEP = errors.New("CEP inválido")
)