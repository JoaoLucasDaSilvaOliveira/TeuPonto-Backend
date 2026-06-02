package exceptions

import "errors"

var (
	ErrInvalidCPF = errors.New("cpf invalido")
)