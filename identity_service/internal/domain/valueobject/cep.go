package valueobject

import (
	"fmt"
	"strings"

	"github.com/paemuri/brdoc"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type CEP string

func NewCEP(rawCEP string, federativeUnit brdoc.FederativeUnit) (CEP, error) {
	trimmedCEP := strings.TrimSpace(rawCEP)

	if trimmedCEP == "" {
		return "", fmt.Errorf("%w: forneça um CEP", exceptions.ErrEmptyString)
	}

	isCEP := brdoc.IsCEP(rawCEP, federativeUnit)

	if !isCEP {
		return "", fmt.Errorf("%w: forneça um CEP", exceptions.ErrInvalidCEP)
	}

	return CEP(trimmedCEP), nil
}