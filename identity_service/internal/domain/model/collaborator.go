package model

import (
	"github.com/google/uuid"
)

type Collaborator struct {
	*Person
	id          uuid.UUID
	idCompany   uuid.UUID
}

func (c *Collaborator) GetID() uuid.UUID {
	return c.id
}

func (c *Collaborator) GetIDCompany() uuid.UUID {
	return c.idCompany
}

func (c *Collaborator) SetIDCompany(idEmpresa uuid.UUID) {
	c.idCompany = idEmpresa
}

func NewCollaborator(person *Person, idEmpresa uuid.UUID) *Collaborator {
	collaborator := new(Collaborator)
	collaborator.id = uuid.New()
	collaborator.Person = person
	collaborator.SetIDCompany(idEmpresa)

	return collaborator
}
