package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type WorkHoursRepository interface {
	SimpleCrudRepository[entity.WorkHours]
	//TODO: create new methods according to the development
}
