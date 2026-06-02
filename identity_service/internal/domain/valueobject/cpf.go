package valueobject

import (
	"strings"

	"github.com/paemuri/brdoc"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type CPF string

func NewCPF(rawCPF string) (CPF, error) {
	//retirar os espaços
	rawCPF = strings.TrimSpace(rawCPF)

	if !brdoc.IsCPF(rawCPF) {
		return "", exceptions.ErrInvalidCPF
	}

	//retorna um cast de CPF, usa o retorno da função que separa somente os dígitos da string cpf
	return CPF(onlyDigits(rawCPF)), nil
}

func (c CPF) String() string {
	return string(c)
}
