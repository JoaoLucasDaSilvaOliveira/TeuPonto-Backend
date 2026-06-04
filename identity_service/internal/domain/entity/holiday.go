package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type Holiday struct {
	*model.Auditable
	id       uuid.UUID
	optional bool
	codIbge  int
	date     *time.Time
}

func (h *Holiday) GetID() uuid.UUID {
	return h.id
}

func (h *Holiday) GetAuditable() *model.Auditable {
	return h.Auditable
}

func (h *Holiday) GetOptional() bool {
	return h.optional
}

func (h *Holiday) SetOptional(optional bool) {
	//altera o updated_at para auditoria
	now := time.Now()
	h.SetUpdatedAt(&now)
	
	h.optional = optional
}

func (h *Holiday) GetCodIbge() int {
	return h.codIbge
}

func (h *Holiday) SetCodIbge(codIbge int) {
	//altera o updated_at para auditoria
	now := time.Now()
	h.SetUpdatedAt(&now)
	
	h.codIbge = codIbge
}

func (h *Holiday) GetDate() *time.Time {
	return h.date
}

func (h *Holiday) SetDate(date *time.Time) error {
	if date == nil {
		return fmt.Errorf("%w: date nula", exceptions.ErrEmptyString)
	}

	h.date = date

	//altera o updated_at para auditoria
	now := time.Now()
	h.SetUpdatedAt(&now)
	
	return nil
}

func NewHoliday(optional bool, codIbge int, date *time.Time) (*Holiday, error) {
	auditable, err := model.NewAuditable()
	if err != nil {
		return nil, err
	}

	if date == nil {
		return nil, fmt.Errorf("%w: date nula", exceptions.ErrEmptyString)
	}

	return &Holiday{
		Auditable: auditable,
		id:        uuid.New(),
		optional:  optional,
		codIbge:   codIbge,
		date:      date,
	}, nil
}
