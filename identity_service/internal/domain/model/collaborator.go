package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Collaborator struct {
	*Person
	id          uuid.UUID
	idEmpresa   uuid.UUID
	password    string
	link_acesso string
}

func (c *Collaborator) GetID() uuid.UUID {
	return c.id
}

func (c *Collaborator) GetIDEmpresa() uuid.UUID {
	return c.idEmpresa
}

func (c *Collaborator) SetIDEmpresa(idEmpresa uuid.UUID) {
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.idEmpresa = idEmpresa
}

func (c *Collaborator) GetPassword() string {
	return c.password
}

func (c *Collaborator) SetPassword(password string) {
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.password = password
}

func (c *Collaborator) GetLinkAcesso() string {
	return c.link_acesso
}

func (c *Collaborator) SetLinkAcesso(linkAcesso string) error {
	trimmedLinkAcesso := strings.TrimSpace(linkAcesso)
	if trimmedLinkAcesso == "" {
		return fmt.Errorf("%w: forneça um link de acesso", exceptions.ErrEmptyString)
	}

	c.link_acesso = trimmedLinkAcesso

	now := time.Now()
	c.SetUpdatedAt(&now)

	return nil
}

func NewCollaborator(person *Person, idEmpresa uuid.UUID, password string, linkAcesso string) (*Collaborator, error) {
	trimmedLinkAcesso := strings.TrimSpace(linkAcesso)
	if trimmedLinkAcesso == "" {
		return nil, fmt.Errorf("%w: forneça um link de acesso", exceptions.ErrEmptyString)
	}

	return &Collaborator{
		Person:      person,
		id:          uuid.New(),
		idEmpresa:   idEmpresa,
		password:    password,
		link_acesso: trimmedLinkAcesso,
	}, nil
}


