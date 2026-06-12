package model

import (
	"fmt"
)

type CalculationMode string

const (
	Direct   CalculationMode = "DIRECT"
	TimeBank CalculationMode = "TIME_BANK"
)

type TimeBankCalculationPolicy struct {
	cycleMonths              int
	entryToleranceMinutes    int
	dailyToleranceMinutes    int
	closingCycleExtraPercent int
}

func (tbcp TimeBankCalculationPolicy) Mode() CalculationMode {
	return TimeBank
}

func (tbcp TimeBankCalculationPolicy) GetCycleMonths() int {
	return tbcp.cycleMonths
}

func (tbcp *TimeBankCalculationPolicy) SetCycleMonths(cycleMonths int) error {
	if err := validateNonNegative("ciclo de meses", cycleMonths); err != nil {
		return err
	}

	tbcp.cycleMonths = cycleMonths
	return nil
}

func (tbcp TimeBankCalculationPolicy) GetEntryToleranceMinutes() int {
	return tbcp.entryToleranceMinutes
}

func (tbcp *TimeBankCalculationPolicy) SetEntryToleranceMinutes(entryToleranceMinutes int) error {
	if err := validateNonNegative("tolerância em minutos", entryToleranceMinutes); err != nil {
		return err
	}

	tbcp.entryToleranceMinutes = entryToleranceMinutes
	return nil
}

func (tbcp TimeBankCalculationPolicy) GetDailyToleranceMinutes() int {
	return tbcp.dailyToleranceMinutes
}

func (tbcp *TimeBankCalculationPolicy) SetDailyToleranceMinutes(dailyToleranceMinutes int) error {
	if err := validateNonNegative("tolerância diária em minutos", dailyToleranceMinutes); err != nil {
		return err
	}

	tbcp.dailyToleranceMinutes = dailyToleranceMinutes
	return nil
}

func (tbcp TimeBankCalculationPolicy) GetClosingCycleExtraPercent() int {
	return tbcp.closingCycleExtraPercent
}

func (tbcp *TimeBankCalculationPolicy) SetClosingCycleExtraPercent(closingCycleExtraPercent int) error {
	if err := validateNonNegative("porcentagem das horas extras", closingCycleExtraPercent); err != nil {
		return err
	}

	tbcp.closingCycleExtraPercent = closingCycleExtraPercent
	return nil
}

func NewDefaultTimeBankCalculationPolicy() (*TimeBankCalculationPolicy, error) {
	return NewTimeBankCalculationPolicy(6, 5, 10, 50)
}

func NewTimeBankCalculationPolicy(
	cycleMonths int,
	entryToleranceMinutes int,
	dailyToleranceMinutes int,
	closingCycleExtraPercent int,
) (*TimeBankCalculationPolicy, error) {
	timeBankCalculationPolicy := new(TimeBankCalculationPolicy)
	
	if err := timeBankCalculationPolicy.SetClosingCycleExtraPercent(closingCycleExtraPercent); err != nil {
		return nil, err
	}
	
	if err := timeBankCalculationPolicy.SetCycleMonths(cycleMonths); err != nil {
		return nil, err
	}
	
	if err := timeBankCalculationPolicy.SetDailyToleranceMinutes(dailyToleranceMinutes); err != nil {
		return nil, err
	}
	
	if err := timeBankCalculationPolicy.SetEntryToleranceMinutes(entryToleranceMinutes); err != nil {
		return nil, err
	}
	
	return timeBankCalculationPolicy, nil
}

type DirectCalculationPolicy struct {
	entryToleranceMinutes                  int
	dailyToleranceMinutes                  int
	weekdayExtraPercentBeforeFirstTwoHours int
	weekdayExtraPercentAfterFirstTwoHours  int
	sundayExtraPercent                     int
	holidayExtraPercent                    int
}

func (dcp DirectCalculationPolicy) Mode() CalculationMode {
	return Direct
}

func (dcp DirectCalculationPolicy) GetEntryToleranceMinutes() int {
	return dcp.entryToleranceMinutes
}

func (dcp *DirectCalculationPolicy) SetEntryToleranceMinutes(entryToleranceMinutes int) error {
	if err := validateNonNegative("tolerâcia em minutos", entryToleranceMinutes); err != nil {
		return err
	}

	dcp.entryToleranceMinutes = entryToleranceMinutes
	return nil
}

func (dcp DirectCalculationPolicy) GetDailyToleranceMinutes() int {
	return dcp.dailyToleranceMinutes
}

func (dcp *DirectCalculationPolicy) SetDailyToleranceMinutes(dailyToleranceMinutes int) error {
	if err := validateNonNegative("tolerâcia diária em minutos", dailyToleranceMinutes); err != nil {
		return err
	}

	dcp.dailyToleranceMinutes = dailyToleranceMinutes
	return nil
}

func (dcp DirectCalculationPolicy) GetWeekdayExtraPercentBeforeFirstTwoHours() int {
	return dcp.weekdayExtraPercentBeforeFirstTwoHours
}

func (dcp *DirectCalculationPolicy) SetWeekdayExtraPercentBeforeFirstTwoHours(weekdayExtraPercentBeforeFirstTwoHours int) error {
	if err := validateNonNegative("porcentagem das horas extras até as 2 primeiras horas extras", weekdayExtraPercentBeforeFirstTwoHours); err != nil {
		return err
	}

	dcp.weekdayExtraPercentBeforeFirstTwoHours = weekdayExtraPercentBeforeFirstTwoHours
	return nil
}

func (dcp DirectCalculationPolicy) GetWeekdayExtraPercentAfterFirstTwoHours() int {
	return dcp.weekdayExtraPercentAfterFirstTwoHours
}

func (dcp *DirectCalculationPolicy) SetWeekdayExtraPercentAfterFirstTwoHours(weekdayExtraPercentAfterFirstTwoHours int) error {
	if err := validateNonNegative("porcentagem das horas extras após as 2 primeiras horas extras", weekdayExtraPercentAfterFirstTwoHours); err != nil {
		return err
	}

	dcp.weekdayExtraPercentAfterFirstTwoHours = weekdayExtraPercentAfterFirstTwoHours
	return nil
}

func (dcp DirectCalculationPolicy) GetSundayExtraPercent() int {
	return dcp.sundayExtraPercent
}

func (dcp *DirectCalculationPolicy) SetSundayExtraPercent(sundayExtraPercent int) error {
	if err := validateNonNegative("porcentagem das horas extras aos domingos", sundayExtraPercent); err != nil {
		return err
	}

	dcp.sundayExtraPercent = sundayExtraPercent
	return nil
}

func (dcp DirectCalculationPolicy) GetHolidayExtraPercent() int {
	return dcp.holidayExtraPercent
}

func (dcp *DirectCalculationPolicy) SetHolidayExtraPercent(holidayExtraPercent int) error {
	if err := validateNonNegative("porcentagem das horas extras em feriados", holidayExtraPercent); err != nil {
		return err
	}

	dcp.holidayExtraPercent = holidayExtraPercent
	return nil
}

func NewDefaultDirectCalculationPolicy() (*DirectCalculationPolicy, error) {
	return NewDirectCalculationPolicy(5, 10, 50, 100, 100, 100)
}

func NewDirectCalculationPolicy(
	entryToleranceMinutes int,
	dailyToleranceMinutes int,
	weekdayExtraPercentBeforeFirstTwoHours int,
	weekdayExtraPercentAfterFirstTwoHours int,
	sundayExtraPercent int,
	holidayExtraPercent int,
) (*DirectCalculationPolicy, error) {
	d := new(DirectCalculationPolicy)

	if err := d.SetDailyToleranceMinutes(dailyToleranceMinutes); err != nil {
		return nil, err
	}
	
	if err := d.SetEntryToleranceMinutes(entryToleranceMinutes); err != nil {
		return nil, err
	}
	
	if err := d.SetHolidayExtraPercent(holidayExtraPercent); err != nil {
		return nil, err
	}
	
	if err := d.SetSundayExtraPercent(sundayExtraPercent); err != nil {
		return nil, err
	}
	
	if err := d.SetWeekdayExtraPercentAfterFirstTwoHours(weekdayExtraPercentAfterFirstTwoHours); err != nil {
		return nil, err
	}
	
	if err := d.SetWeekdayExtraPercentBeforeFirstTwoHours(weekdayExtraPercentBeforeFirstTwoHours); err != nil {
		return nil, err
	}
	
	return &DirectCalculationPolicy{
		entryToleranceMinutes:                  entryToleranceMinutes,
		dailyToleranceMinutes:                  dailyToleranceMinutes,
		weekdayExtraPercentBeforeFirstTwoHours: weekdayExtraPercentBeforeFirstTwoHours,
		weekdayExtraPercentAfterFirstTwoHours:  weekdayExtraPercentAfterFirstTwoHours,
		sundayExtraPercent:                     sundayExtraPercent,
		holidayExtraPercent:                    holidayExtraPercent,
	}, nil
}

func validateNonNegative(fieldName string, value int) error {
	if value < 0 {
		return fmt.Errorf("%s não pode ser negativo", fieldName)
	}

	return nil
}

type CalculationPolicy interface {
	Mode() CalculationMode
}
