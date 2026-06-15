package valueobject

import (
	"strings"

	"github.com/paemuri/brdoc"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type CPF string

func NewCPF(rawCPF string) (CPF, error) {
	//remove spaces
	rawCPF = strings.TrimSpace(rawCPF)

	if !brdoc.IsCPF(rawCPF) {
		return "", exceptions.ErrInvalidCPF
	}

	// returns a cast of the CPF, using the return value of the function that extracts only the digits from the CPF string.
	return CPF(onlyDigits(rawCPF)), nil
}

func (c CPF) String() string {
	return string(c)
}
