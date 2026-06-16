package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type WorkerRepository interface {
	SimpleCrudRepository[entity.Worker]
	//TODO: create new methods according to the development
}
