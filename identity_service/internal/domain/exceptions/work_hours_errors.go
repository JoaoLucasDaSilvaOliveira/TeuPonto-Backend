package exceptions

import "errors"

var (
	ErrInvalidTime = errors.New("horário inválido")
	ErrUnprocessableEntryTime = errors.New("horário de entrada é depois ou igual ao horário de saída")
	ErrUnprocessableExitTime = errors.New("horário de saída é anterior ou igual ao horário de entrada")
)