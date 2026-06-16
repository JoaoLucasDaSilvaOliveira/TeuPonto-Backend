package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type OutsourcedRepository interface {
	SimpleCrudRepository[entity.Outsourced]
	//TODO: create new methods according to the development
}
