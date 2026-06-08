package entity

import (
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type Manager struct {
	*model.Collaborator
}

func NewManager(collaborator *model.Collaborator) (*Manager, error) {
	return &Manager{
		Collaborator: collaborator,
	}, nil
}
