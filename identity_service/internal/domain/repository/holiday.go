package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type HolidayRepository interface {
	SimpleCrudRepository[entity.Holiday]
	//TODO: create new methods according to the development
}
