package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type AddressRepository interface {
	SimpleCrudRepository[entity.Address]
	//TODO: create new methods according to the development
}