package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
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
	
	return nil
}

func (w *WorkShift) GetDescription() string {
	return w.description
}

func (w *WorkShift) SetDescription(description string) {
	w.description = strings.TrimSpace(description)
}

func (w *WorkShift) GetShiftType() ShiftType {
	return w.shiftType
}

func (w *WorkShift) SetShiftType(shiftType ShiftType) {
	w.shiftType = shiftType
}

func (w *WorkShift) GetWorkHours() *WorkHours {
	return w.workHours
}

func (w *WorkShift) SetWorkHours(workHours *WorkHours) {
	w.workHours = workHours
}

func NewWorkShift(
	name string,
	description string,
	shiftType ShiftType,
	workHours *WorkHours,
) (*WorkShift, error) {
	workShift := new(WorkShift)

	workShift.id = uuid.New()
	workShift.SetDescription(description)
	workShift.SetShiftType(shiftType)
	workShift.SetWorkHours(workHours)	
	
	if err := workShift.SetName(name); err != nil {
		return nil, err
	}

	return workShift, nil
}
