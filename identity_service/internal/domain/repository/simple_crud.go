package repository

import (
	"github.com/google/uuid"
)

type SimpleCrudRepository[T SimpleCrudEntity] interface {
	Create(entity *T) (uuid.UUID, error)
	GetByID(id uuid.UUID) (*T, error)
	FindAll(filter map[string]any) ([]*T, error)
	Update(newEntity *T) error
	Delete(id uuid.UUID) error
	ExistsById(id uuid.UUID) (bool, error)
}
