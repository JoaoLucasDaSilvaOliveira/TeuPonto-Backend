package valueobject

import (
	"fmt"
	"net/mail"
	"strings"

	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Email string

func NewEmail(rawEmail string) (Email, error) {
	rawEmail = strings.TrimSpace(rawEmail)

	if rawEmail == "" {
		return "", fmt.Errorf("%w: email fornecido está vazio", exceptions.ErrInvalidEmail)
	}

	_, err := mail.ParseAddress(rawEmail)

	if err != nil {
		return "", exceptions.ErrInvalidEmail
	}

	return Email(strings.ToLower(rawEmail)), nil
}