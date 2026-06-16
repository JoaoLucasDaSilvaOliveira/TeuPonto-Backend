package repository

import "teuponto.com.br/backend/identity_service/internal/domain/entity"

type CompanyRepository interface {
	SimpleCrudRepository[entity.Company]
	//TODO: create new methods according to the development
}
