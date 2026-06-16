package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type WorkShiftRepository interface {
	SimpleCrudRepository[entity.WorkShift]
	//TODO: create new methods according to the development
}
