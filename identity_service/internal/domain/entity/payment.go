package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Payment struct {
	id              uuid.UUID
	idPlan          uuid.UUID
	startContract   *time.Time
	dateLastPayment *time.Time
}

func (p *Payment) GetID() uuid.UUID {
	return p.id
}

func (p *Payment) GetIDPlan() uuid.UUID {
	return p.idPlan
}

func (p *Payment) SetIDPlan(idPlan uuid.UUID) {
	p.idPlan = idPlan
}

func (p *Payment) GetStartContract() *time.Time {
	return p.startContract
}

func (p *Payment) SetStartContract(startContract *time.Time) error {
	if startContract == nil {
		return fmt.Errorf("%w: startContract nulo", exceptions.ErrEmptyString)
	}

	p.startContract = startContract
	
	return nil
}

func (p *Payment) GetDateLastPayment() *time.Time {
	return p.dateLastPayment
}

func (p *Payment) SetDateLastPayment(dateLastPayment *time.Time) error {
	if dateLastPayment == nil {
		return fmt.Errorf("%w: dateLastPayment nulo", exceptions.ErrEmptyString)
	}

	p.dateLastPayment = dateLastPayment

	return nil
}

func NewPayment(idPlan uuid.UUID, startContract *time.Time, dateLastPayment *time.Time) (*Payment, error) {
	payment := new(Payment)

	payment.id = uuid.New()
	payment.SetIDPlan(idPlan)

	if err := payment.SetDateLastPayment(dateLastPayment); err != nil {
		return nil, err
	}

	if err := payment.SetStartContract(startContract); err != nil {
		return nil, err
	}	

	return payment, nil
}
