package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type ManagerRepository interface {
	SimpleCrudRepository[entity.Manager]
	//TODO: create new methods according to the development
}
