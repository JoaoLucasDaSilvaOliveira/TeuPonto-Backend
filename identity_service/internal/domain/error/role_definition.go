package error

import "errors"

var (
	ErrInvalidRole = errors.New("role inválida")
	ErrPermissionNotAllowed = errors.New("permissão não é compatível com a role")
)