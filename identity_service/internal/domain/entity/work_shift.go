package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
)

type ShiftType string

const (
	shift_6X1    ShiftType = "6X1"
	shift_5X2    ShiftType = "5X2"
	shift_4X3    ShiftType = "4X3"
	shift_12X36  ShiftType = "12X36"
	shift_custom ShiftType = "custom_shift"
)

type WorkShift struct {
	*model.Auditable
	id          uuid.UUID
	name        string
	description string
	shiftType   ShiftType
	workHours   *WorkHours
}

func (w *WorkShift) GetID() uuid.UUID {
	return w.id
}

func (w *WorkShift) GetName() string {
	return w.name
}

func (w *WorkShift) SetName(name string) error {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)
	}

	w.name = trimmedName

	now := time.Now()
	w.SetUpdatedAt(&now)

	return nil
}

func (w *WorkShift) GetDescription() string {
	return w.description
}

func (w *WorkShift) SetDescription(description string) {
	w.description = strings.TrimSpace(description)

	now := time.Now()
	w.SetUpdatedAt(&now)
}

func (w *WorkShift) GetShiftType() ShiftType {
	return w.shiftType
}

func (w *WorkShift) SetShiftType(shiftType ShiftType) {
	w.shiftType = shiftType

	now := time.Now()
	w.SetUpdatedAt(&now)
}

func (w *WorkShift) GetWorkHours() *WorkHours {
	return w.workHours
}

func (w *WorkShift) SetWorkHours(workHours *WorkHours) {
	w.workHours = workHours

	now := time.Now()
	w.SetUpdatedAt(&now)
}

func NewWorkShift(
	auditable *model.Auditable,
	name string,
	description string,
	shiftType ShiftType,
	workHours *WorkHours,
) (*WorkShift, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)
	}

	return &WorkShift{
		Auditable:   auditable,
		id:          uuid.New(),
		name:        trimmedName,
		description: strings.TrimSpace(description),
		shiftType:   shiftType,
		workHours:   workHours,
	}, nil
}
