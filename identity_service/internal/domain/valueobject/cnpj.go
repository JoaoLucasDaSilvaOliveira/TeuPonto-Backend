package valueobject

import (
	"strings"

	"github.com/paemuri/brdoc"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type CNPJ string

func NewCNPJ(rawCNPJ string) (CNPJ, error) {
	rawCNPJ = strings.TrimSpace(rawCNPJ)

	if !brdoc.IsCNPJ(rawCNPJ) {
		return "", exceptions.ErrInvalidCNPJ
	}

	return CNPJ(onlyDigits(rawCNPJ)), nil
}

func (c CNPJ) String() string {
	return string(c)
}