package usecase

import (
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type ReportUseCase interface {
	GenerateExpenditureReport(startDate time.Time, endDate time.Time) (*model.ExpenditureReport, error)
}

type reportUseCase struct {
	dailyExpenditureRepo repository.DailyExpenditureRepository
	userUseCase
}

func NewReportUseCase(dailyExpenditureRepo repository.DailyExpenditureRepository) ReportUseCase {
	return &reportUseCase{
		dailyExpenditureRepo: dailyExpenditureRepo,
	}
}

func (uc *reportUseCase) GenerateExpenditureReport(startDate time.Time, endDate time.Time) (*model.ExpenditureReport, error) {
	total, err := uc.dailyExpenditureRepo.GetTotalExpenditureByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	expenditures, err := uc.dailyExpenditureRepo.GetExpendituresByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	report := &model.ExpenditureReport{
		StartDate:        startDate,
		EndDate:          endDate,
		TotalExpenditure: total,
		Expenditures:     []*model.DailyExpenditureReport{},
	}

	for _, expenditure := range expenditures {
		dailyReport := &model.DailyExpenditureReport{
			ID:          expenditure.ID,
			UserID:      expenditure.UserID,
			Username:    expenditure.Username,
			Amount:      expenditure.Amount,
			Description: expenditure.Description,
			CreatedAt:   expenditure.CreatedAt,
			UpdatedAt:   expenditure.UpdatedAt,
			Date:        expenditure.Date,
		}

		report.Expenditures = append(report.Expenditures, dailyReport)
	}

	return report, nil
}


