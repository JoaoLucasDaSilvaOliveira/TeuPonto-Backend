package exceptions

import "errors"

var (
	ErrInvalidCNPJ = errors.New("cnpj invalido")
)