package entity

import (
	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type Supervisor struct {
	*model.Collaborator
	idManager uuid.UUID
}

func (s *Supervisor) GetIDManager() uuid.UUID {
	return s.idManager
}

func (s *Supervisor) SetIDManager(idGestor uuid.UUID) {
	s.idManager = idGestor
}

func NewSupervisor(collaborator *model.Collaborator, idGestor uuid.UUID) (*Supervisor, error) {
	supervisor := new(Supervisor)

	supervisor.SetIDManager(idGestor)
	supervisor.Collaborator = collaborator
	
	return supervisor, nil
}
