package entity

import (
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type Worker struct {
	*model.Collaborator
	idSupervisor uuid.UUID
}

func (w *Worker) GetIDSupervisor() uuid.UUID {
	return w.idSupervisor
}

func (w *Worker) SetIDSupervisor(idSupervisor uuid.UUID) {
	now := time.Now()
	w.SetUpdatedAt(&now)

	w.idSupervisor = idSupervisor
}

func NewWorker(collaborator *model.Collaborator, idSupervisor uuid.UUID) (*Worker, error) {
	return &Worker{
		Collaborator: collaborator,
		idSupervisor: idSupervisor,
	}, nil
}
