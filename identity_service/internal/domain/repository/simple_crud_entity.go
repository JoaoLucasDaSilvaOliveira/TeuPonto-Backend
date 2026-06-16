package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type SimpleCrudEntity interface {
	entity.Address |
		entity.Company |
		entity.Holiday |
		entity.Manager |
		entity.Outsourced |
		entity.Payment |
		entity.Supervisor |
		entity.WorkHours |
		entity.WorkShift |
		entity.Worker
}
