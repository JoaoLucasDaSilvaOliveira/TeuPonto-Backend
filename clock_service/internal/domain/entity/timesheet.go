//idea

package entity

import (
	"strings"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type TimeAdjustment struct {
	idResponsible          uuid.UUID
	newClockInOrOut        civil.Time
	manualAdjustment       bool
	adjustementDescription string
}

func (t *TimeAdjustment) GetIDResponsible() uuid.UUID {
	return t.idResponsible
}

func (t *TimeAdjustment) GetNewClockInOrOut() civil.Time {
	return t.newClockInOrOut
}

func (t *TimeAdjustment) GetManualAdjustment() bool {
	return t.manualAdjustment
}

func (t *TimeAdjustment) GetAdjustmentDescription() string {
	return t.adjustementDescription
}

func NewTimeAdjustment(
	idResponsible uuid.UUID,
	newClockInOrOut civil.Time,
	manualAdjustment bool,
	adjustmentDescription string,
) *TimeAdjustment {
	timeAdjustment := new(TimeAdjustment)

	timeAdjustment.idResponsible = idResponsible 
	timeAdjustment.newClockInOrOut = newClockInOrOut
	timeAdjustment.manualAdjustment = manualAdjustment

	timeAdjustment.adjustementDescription = strings.TrimSpace(adjustmentDescription)

	return timeAdjustment
}

//--------------------------------------------

type TimeClock struct {
	id            uuid.UUID
	idEmployee    uuid.UUID
	clockInOrOut  civil.Time
	timeAjustment *[]TimeAdjustment
}

func (t *TimeClock) GetID() uuid.UUID {
	return t.id
}

func (t *TimeClock) GetIDEmployee() uuid.UUID {
	return t.idEmployee
}

func (t *TimeClock) SetIDEmployee(idEmployee uuid.UUID) {
	t.idEmployee = idEmployee
}

func (t *TimeClock) GetClockInOrOut() civil.Time {
	return t.clockInOrOut
}

func (t *TimeClock) SetClockInOrOut(clockInOrOut civil.Time) {
	t.clockInOrOut = clockInOrOut
}

func (t *TimeClock) GetTimeAdjustment() *TimeAdjustment {
	return t.timeAjustment
}

func (t *TimeClock) SetTimeAdjustment(timeAdjustment *TimeAdjustment) {
	t.timeAjustment = timeAdjustment
}

func NewTimeClock(
	idEmployee uuid.UUID,
	clockInOrOut civil.Time,
	timeAdjustment *TimeAdjustment,
) *TimeClock {
	timeClock := new(TimeClock)

	timeClock.id = uuid.New()
	timeClock.SetIDEmployee(idEmployee)
	timeClock.SetClockInOrOut(clockInOrOut)
	timeClock.SetTimeAdjustment(timeAdjustment)

	return timeClock
}

//--------------------------------------------

type Timesheet struct {
	id         uuid.UUID
	idCompany  uuid.UUID
	timeClocks []*TimeClock
}

func (t *Timesheet) GetID() uuid.UUID {
	return t.id
}

func (t *Timesheet) GetIDCompany() uuid.UUID {
	return t.idCompany
}

func (t *Timesheet) SetIDCompany(idCompany uuid.UUID) {
	t.idCompany = idCompany
}

func (t *Timesheet) GetTimeClocks() []*TimeClock {
	return t.timeClocks
}

func (t *Timesheet) SetTimeClocks(timeClocks []*TimeClock) {
	t.timeClocks = timeClocks
}

func (t *Timesheet) AddTimeClock(timeClock *TimeClock) {
	t.timeClocks = append(t.timeClocks, timeClock)
}

func NewTimesheet(idCompany uuid.UUID, timeClocks []*TimeClock) *Timesheet {
	timesheet := new(Timesheet)

	timesheet.id = uuid.New()
	timesheet.SetIDCompany(idCompany)
	timesheet.SetTimeClocks(timeClocks)

	return timesheet
}
