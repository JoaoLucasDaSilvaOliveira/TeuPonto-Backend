package error

import "errors"

var (
	ErrLoadLocation = errors.New("erro ao obter a localização do sistema")
)