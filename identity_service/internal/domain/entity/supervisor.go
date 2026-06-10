package entity

import (
	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type Supervisor struct {
	*model.Collaborator
	idGestor uuid.UUID
}

func (s *Supervisor) GetIDGestor() uuid.UUID {
	return s.idGestor
}

func (s *Supervisor) SetIDGestor(idGestor uuid.UUID) {
	s.idGestor = idGestor
}

func NewSupervisor(collaborator *model.Collaborator, idGestor uuid.UUID) (*Supervisor, error) {
	return &Supervisor{
		Collaborator: collaborator,
		idGestor:     idGestor,
	}, nil
}
