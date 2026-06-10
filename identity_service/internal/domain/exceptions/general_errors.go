package exceptions

import "errors"

var (
	ErrLoadLocation = errors.New("erro ao obter a localização do sistema")
	ErrEmptyString = errors.New("dado recebido está em branco")
	ErrEmptyCalculationPolicy = errors.New("política de cálculo padrão não pode ser vazia")
)