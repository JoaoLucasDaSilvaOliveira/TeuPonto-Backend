package entity

import (
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
	w.idSupervisor = idSupervisor
}

func NewWorker(collaborator *model.Collaborator, idSupervisor uuid.UUID) (*Worker, error) {
	worker := new(Worker)

	worker.SetIDSupervisor(idSupervisor)
	worker.Collaborator = collaborator
	
	return worker, nil
}
