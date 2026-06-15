package entity

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type shift struct {
	entryTime civil.Time // não pode ser menor que a saída
	exitTime  civil.Time // não pode ser maior que a entrada
}

func (s shift) GetEntryTime() civil.Time {
	return s.entryTime
}

func (s shift) WithEntryTime(entryTime civil.Time) error {
	//verifica se é depois que a saída
	if entryTime.After(s.exitTime) || entryTime == s.entryTime {
		return exceptions.ErrUnprocessableEntryTime
	}
	s.entryTime = entryTime

	return nil
}

func (s shift) GetExitTime() civil.Time {
	return s.exitTime
}

func (s shift) WithExitTime(exitTime civil.Time) error {
	//verifica se é antes da entrada
	if exitTime.Before(s.entryTime) || exitTime == s.exitTime {
		return exceptions.ErrUnprocessableExitTime
	}
	s.exitTime = exitTime

	return nil
}

func NewShift (entryTime civil.Time, exitTime civil.Time) (*shift, error) {
	shift := new(shift)
	
	if err := shift.WithEntryTime(entryTime); err != nil {
		return nil, err
	}

	if err := shift.WithExitTime(exitTime); err != nil {
		return nil, err
	}

	return shift, nil
}

//--------------------------------------------

type WorkHours struct {
	id          uuid.UUID
	name        string
	description string
	workShift   []*shift
	workload    time.Duration
}	

func (w *WorkHours) GetID() uuid.UUID {
	return w.id
}

func (w *WorkHours) GetName() string {
	return w.name
}

func (w *WorkHours) SetName(name string) error {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return fmt.Errorf("%w: forneça um nome", exceptions.ErrEmptyString)
	}

	w.name = trimmedName

	return nil
}

func (w *WorkHours) GetDescription() string {
	return w.description
}

func (w *WorkHours) SetDescription(description string) {
	w.description = strings.TrimSpace(description)
}

func (w *WorkHours) GetWorkShift() []*shift {
	return w.workShift
}

func (w *WorkHours) SetWorkShift(workShift []*shift) {
	w.workShift = workShift
	w.workload = calculateWorkload(w.workShift)
}

func (w *WorkHours) AddShift(newShift *shift) error {
	// check for duplicated
	// use ContainsFunc to compare the value of hours,
	// the simple Contains would compare the pointer's memory addresses  
	if slices.ContainsFunc(w.workShift, func(s *shift) bool {
		return s.entryTime == newShift.entryTime && s.exitTime == newShift.exitTime
	}) {
		return fmt.Errorf("%w: turno de trabalho já cadastrado", exceptions.ErrDuplicatedElement)
	}

	// check temporal logic
	if len(w.workShift) > 0 {
		lastShift := w.workShift[len(w.workShift)-1]

		// the new entry time cannot be before (or in the exact minute) of the last exit
		if newShift.entryTime.Before(lastShift.exitTime) || newShift.entryTime == lastShift.exitTime {
			// O ideal seria criar um exceptions.ErrShiftOverlap no seu pacote de erros
			return fmt.Errorf("%w: conflito de horários: a entrada (%v) não pode ser anterior ou igual à saída do turno anterior (%v)", exceptions.ErrInvalidTime, newShift.entryTime, lastShift.exitTime)
		}
	}
	
	w.workShift = append(w.workShift, newShift)
	w.workload = calculateWorkload(w.workShift)

	return nil
}

func (w *WorkHours) GetWorkload() time.Duration {
	return w.workload
}

func NewWorkHours(name string, description string, workShift []*shift) (*WorkHours, error) {
	workHours := new(WorkHours)

	workHours.id = uuid.New()

	workHours.SetDescription(description)
	workHours.SetWorkShift(workShift)

	if err := workHours.SetName(name); err != nil {
		return nil, err
	}

	workHours.workload = calculateWorkload(workHours.GetWorkShift())

	return workHours, nil
}

func calculateWorkload(workShift []*shift) time.Duration {
	var workload time.Duration

	for _, workShiftItem := range workShift {
		workload += shiftDuration(workShiftItem)
	}

	return workload
}

func shiftDuration(workShift *shift) time.Duration {
	entry := time.Date(2000, time.January, 1, workShift.entryTime.Hour, workShift.entryTime.Minute, workShift.entryTime.Second, 0, time.UTC)

	// Dia padrão é 1
	exitDay := 1
	// Se a hora de saída for menor que a de entrada, significa que passou da meia-noite (dia seguinte)
	if workShift.exitTime.Before(workShift.entryTime) {
		exitDay = 2
	}

	exit := time.Date(2000, time.January, exitDay, workShift.exitTime.Hour, workShift.exitTime.Minute, workShift.exitTime.Second, 0, time.UTC)

	return exit.Sub(entry)
}