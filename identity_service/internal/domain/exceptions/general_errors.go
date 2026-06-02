package exceptions

import "errors"

var (
	ErrLoadLocation = errors.New("erro ao obter a localização do sistema")
	ErrEmptyName = errors.New("nome recebido está em branco, forneça um nome")
)