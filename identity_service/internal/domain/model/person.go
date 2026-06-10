package model

import (
	"fmt"
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
		return fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)
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
	person := new(Person)
	
	if err := person.SetCPF(rawCpf); err != nil {
		return nil, err
	}
	
	if err := person.SetEmail(rawEmail); err != nil {
		return nil, err
	}
	
	if err := person.SetName(name); err != nil {
		return nil, err
	}
	
	person.User = user

	return person, nil
}
