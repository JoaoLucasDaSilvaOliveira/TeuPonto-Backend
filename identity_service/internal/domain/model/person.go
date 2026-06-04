package model

import (
	"fmt"
	"strings"
	"time"

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
		return fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)
	}
	
	p.name = trimmedName

	//altera o updated_at para auditoria
	now := time.Now()
	p.SetUpdatedAt(&now)

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

	//altera o updated_at para auditoria
	now := time.Now()
	p.SetUpdatedAt(&now)
	
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

	//altera o updated_at para auditoria
	now := time.Now()
	p.SetUpdatedAt(&now)
	
	return nil
}

func NewPerson(user *User, name string, rawCpf string, rawEmail string) (*Person, error) {
	trimmedName := strings.TrimSpace(name)

	if trimmedName == "" {
		return nil, fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)

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
