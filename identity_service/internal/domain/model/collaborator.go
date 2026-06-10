package model

import (
	"github.com/google/uuid"
)

type Collaborator struct {
	*Person
	id          uuid.UUID
	idEmpresa   uuid.UUID
}

func (c *Collaborator) GetID() uuid.UUID {
	return c.id
}

func (c *Collaborator) GetIDEmpresa() uuid.UUID {
	return c.idEmpresa
}

func (c *Collaborator) SetIDEmpresa(idEmpresa uuid.UUID) {
	c.idEmpresa = idEmpresa
}

func NewCollaborator(person *Person, idEmpresa uuid.UUID) *Collaborator {
	collaborator := new(Collaborator)
	collaborator.id = uuid.New()
	collaborator.Person = person
	collaborator.SetIDEmpresa(idEmpresa)

	return collaborator
}
