package usecase

import (
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type DailyExpenditureUseCase interface {
	CreateDailyExpenditure(expenditure *model.DailyExpenditure) error
	UpdateDailyExpenditure(expenditure *model.DailyExpenditure) error
	GetDailyExpenditureByID(id string) (*model.DailyExpenditure, error)
	GetAllDailyExpenditures() ([]*model.DailyExpenditure, error)
	DeleteDailyExpenditure(id string) error
	GetTotalExpenditureByDateRange(startDate time.Time, endDate time.Time) (float64, error)
}

type dailyExpenditureUseCase struct {
	dailyExpenditureRepo repository.DailyExpenditureRepository
}

func NewDailyExpenditureUseCase(deRepo repository.DailyExpenditureRepository) DailyExpenditureUseCase {
	return &dailyExpenditureUseCase{
		dailyExpenditureRepo: deRepo,
	}
}

func (uc *dailyExpenditureUseCase) CreateDailyExpenditure(expenditure *model.DailyExpenditure) error {
	// Perform any business logic or validation before creating the daily expenditure
	// ...

	return uc.dailyExpenditureRepo.CreateDailyExpenditure(expenditure)
}

func (uc *dailyExpenditureUseCase) UpdateDailyExpenditure(expenditure *model.DailyExpenditure) error {
	// Perform any business logic or validation before updating the daily expenditure
	// ...

	return uc.dailyExpenditureRepo.UpdateDailyExpenditure(expenditure)
}

func (uc *dailyExpenditureUseCase) GetDailyExpenditureByID(id string) (*model.DailyExpenditure, error) {
	return uc.dailyExpenditureRepo.GetDailyExpenditureByID(id)
}

func (uc *dailyExpenditureUseCase) GetAllDailyExpenditures() ([]*model.DailyExpenditure, error) {
	return uc.dailyExpenditureRepo.GetAllDailyExpenditures()
}

func (uc *dailyExpenditureUseCase) DeleteDailyExpenditure(id string) error {
	return uc.dailyExpenditureRepo.DeleteDailyExpenditure(id)
}

func (uc *dailyExpenditureUseCase) GetTotalExpenditureByDateRange(startDate time.Time, endDate time.Time) (float64, error) {
	return uc.dailyExpenditureRepo.GetTotalExpenditureByDateRange(startDate, endDate)
}
