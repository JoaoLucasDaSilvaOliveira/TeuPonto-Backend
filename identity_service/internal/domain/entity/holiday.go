package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Holiday struct {
	id       uuid.UUID
	optional bool
	codIbge  int
	date     *time.Time
}

func (h *Holiday) GetID() uuid.UUID {
	return h.id
}

func (h *Holiday) GetOptional() bool {
	return h.optional
}

func (h *Holiday) SetOptional(optional bool) {
	h.optional = optional
}

func (h *Holiday) GetCodIbge() int {
	return h.codIbge
}

func (h *Holiday) SetCodIbge(codIbge int) {
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

	return nil
}

func NewHoliday(optional bool, codIbge int, date *time.Time) (*Holiday, error) {
	holiday := new(Holiday)

	holiday.id = uuid.New()

	holiday.SetCodIbge(codIbge)
	holiday.SetOptional(optional)

	if err := holiday.SetDate(date); err != nil {
		return nil, err
	}

	return holiday, nil
}
