package usecase

import (
	"fmt"
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
	GenerateNotaNumber() (string, error)
}

type dailyExpenditureUseCase struct {
	dailyExpenditureRepo repository.DailyExpenditureRepository
	userRepo repository.UserRepository
}

func NewDailyExpenditureUseCase(deRepo repository.DailyExpenditureRepository, userRepo repository.UserRepository) DailyExpenditureUseCase {
	return &dailyExpenditureUseCase{
		dailyExpenditureRepo: deRepo,
		userRepo: userRepo,
	}
}

func (uc *dailyExpenditureUseCase) CreateDailyExpenditure(expenditure *model.DailyExpenditure) error {
	// Perform any business logic or validation before creating the daily expenditure
	// ...
	isExist, err :=uc.userRepo.GetUserByID(expenditure.UserID)
	if err != nil {
		return  err
	}
	if isExist==nil {
		return fmt.Errorf("silahkan login ulang")
	}
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

func (uc *dailyExpenditureUseCase) GenerateNotaNumber() (string, error) {
	now := time.Now().Format("20060102")
	lastNotaNumber, err := uc.dailyExpenditureRepo.GetLastNotaNumber(now)
	if err != nil {
		return "", err
	}

	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")

	noteNumber := fmt.Sprintf("DE-%s%s%s%04d", year, month, day, lastNotaNumber)

	return noteNumber, nil
}
