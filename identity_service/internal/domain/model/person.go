package model

import (
	"strings"

	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type Person struct {
	*User //embedded
	name  string
	cpf   valueobject.CPF
	email valueobject.Email
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) error {
	trimmedName := strings.TrimSpace(name)

	if trimmedName == "" {
		return exceptions.ErrEmptyName
	}

	p.name = trimmedName
	return nil
}

func (p *Person) GetCPF() valueobject.CPF {
	return p.cpf
}

func (p *Person) SetCPF(rawCpf string) error {
	cpf, err := valueobject.NewCPF(rawCpf)
	if err != nil {
		return err
	}

	p.cpf = cpf
	return nil
}

func (p *Person) GetEmail() valueobject.Email {
	return p.email
}

func (p *Person) SetEmail(rawEmail string) error {
	email, err := valueobject.NewEmail(rawEmail)
	if err != nil {
		return err
	}

	p.email = email
	return nil
}

func NewPerson(user *User, name string, rawCpf string, rawEmail string) (*Person, error) {
	trimmedName := strings.TrimSpace(name)

	if trimmedName == "" {
		return nil, exceptions.ErrEmptyName
	}

	cpf, err := valueobject.NewCPF(rawCpf)

	if err != nil {
		return nil, err
	}

	email, err := valueobject.NewEmail(rawEmail)

	if err != nil {
		return nil, err
	}

	return &Person{
		name:  trimmedName,
		cpf:   cpf,
		email: email,
		User:  user,
	}, nil
}
