package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type SupervisorRepository interface {
	SimpleCrudRepository[entity.Supervisor]
	//TODO: create new methods according to the development
}
