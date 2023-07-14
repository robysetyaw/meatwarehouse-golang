package model

import "time"

type ExpenditureReport struct {
	StartDate        time.Time
	EndDate          time.Time
	TotalExpenditure float64
	Expenditures     []*DailyExpenditure
}
