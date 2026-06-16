package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type PaymentRepository interface {
	SimpleCrudRepository[entity.Payment]
	//TODO: create new methods according to the development
}
