package entity

import (
	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type Outsourced struct {
	*model.User
	cnpj           valueobject.CNPJ
}

func (o *Outsourced) GetID() uuid.UUID {
	return o.User.GetUserID()
}	

func (o *Outsourced) GetCNPJ() valueobject.CNPJ {
	return o.cnpj
}

func (o *Outsourced) SetCNPJ(rawCNPJ string) error {
	cnpj, err := valueobject.NewCNPJ(rawCNPJ)
	if err != nil {
		return err
	}

	o.cnpj = cnpj

	return nil
}

func NewOutsourced(user *model.User, rawCNPJ string, accessPassword string) (*Outsourced, error) {
	outsourced := new(Outsourced)

	outsourced.User = user

	if err := outsourced.SetCNPJ(rawCNPJ); err != nil {
		return nil, err
	}

	return outsourced, nil
}
