package entity

import (
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type Outsourced struct {
	*model.User
	cnpj           valueobject.CNPJ
	accessPassword string
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

	now := time.Now()
	o.SetUpdatedAt(&now)

	return nil
}

func (o *Outsourced) GetAccessPassword() string {
	return o.accessPassword
}

func (o *Outsourced) SetAccessPassword(accessPassword string) {
	now := time.Now()
	o.SetUpdatedAt(&now)

	o.accessPassword = accessPassword
}

func NewOutsourced(user *model.User, rawCNPJ string, accessPassword string) (*Outsourced, error) {
	cnpj, err := valueobject.NewCNPJ(rawCNPJ)
	if err != nil {
		return nil, err
	}

	return &Outsourced{
		User:           user,
		cnpj:           cnpj,
		accessPassword: accessPassword,
	}, nil
}
