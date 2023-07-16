package usecase

import (
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type StockMovementUseCase interface {
	GenerateStockMovementReport(startDate time.Time,endDate time.Time,) ([]*model.StockMovementReport, error)
	// Tambahkan fungsi lain yang dibutuhkan untuk use case laporan pergerakan stok
}

type stockMovementUseCase struct {
	meatRepo          repository.MeatRepository
	transactionRepo   repository.TransactionRepository
}

func NewStockMovementUseCase(meatRepo repository.MeatRepository,transactionRepo repository.TransactionRepository,) StockMovementUseCase {
	return &stockMovementUseCase{
		meatRepo:          meatRepo,
		transactionRepo:   transactionRepo,
	}
}

func (uc *stockMovementUseCase) GenerateStockMovementReport(startDate time.Time,endDate time.Time,) ([]*model.StockMovementReport, error) {
	meats, err := uc.meatRepo.GetAllMeats()
	if err != nil {
		return nil, err
	}

	stockMovementReports := make([]*model.StockMovementReport, 0)

	for _, meat := range meats {
		totalStockIn, err := uc.calculateTotalStockIn(startDate,endDate, meat.ID)
		if err != nil {
			return nil, err
		}

		totalStockOut, err := uc.calculateTotalStockOut(startDate,endDate,meat.ID)
		if err != nil {
			return nil, err
		}

		stockMovement := totalStockIn - totalStockOut

		stockMovementReport := &model.StockMovementReport{
			MeatID:        meat.ID,
			MeatName:      meat.Name,
			TotalStockIn:  totalStockIn,
			TotalStockOut: totalStockOut,
			StockMovement: stockMovement,
		}

		stockMovementReports = append(stockMovementReports, stockMovementReport)
	}

	return stockMovementReports, nil
}

func (uc *stockMovementUseCase) calculateTotalStockIn(startDate time.Time,endDate time.Time,meatID string) (float64, error) {

	transactions, err := uc.transactionRepo.GetTransactionsByDateAndType(startDate, endDate , "in")
	if err != nil {
		return 0, err
	}

	totalStockIn := 0.0

	for _, transaction := range transactions {

		for _, detail := range transaction.TransactionDetails {
			if detail.MeatID == meatID {
				totalStockIn += detail.Qty
			}
		}
	}


	return totalStockIn, nil
}

func (uc *stockMovementUseCase) calculateTotalStockOut(startDate time.Time,endDate time.Time,meatID string) (float64, error) {

	transactions, err := uc.transactionRepo.GetTransactionsByDateAndType(startDate, endDate, "out")
	if err != nil {
		return 0, err
	}

	totalStockOut := 0.0
	for _, transaction := range transactions {
		for _, detail := range transaction.TransactionDetails {
			if detail.MeatID == meatID {
				totalStockOut += detail.Qty
			}
		}
	}

	return totalStockOut, nil
}
