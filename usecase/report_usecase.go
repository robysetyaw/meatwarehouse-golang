package usecase

import (
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type ReportUseCase interface {
	GenerateExpenditureReport(startDate time.Time, endDate time.Time) (*model.ExpenditureReport, error)
	GenerateReport(startDate time.Time, endDate time.Time) (*model.TransactionReport, error)
	GenerateSalesReport(startDate time.Time, endDate time.Time) (*model.SalesReportOut, error)
	GenerateReceiptReport(startDate time.Time, endDate time.Time) (*model.ReceiptReport, error)
}

type reportUseCase struct {
	transactionRepo      repository.TransactionRepository
	dailyExpenditureRepo repository.DailyExpenditureRepository
	userUseCase
}

func NewReportUseCase(dailyExpenditureRepo repository.DailyExpenditureRepository, transactionRepo repository.TransactionRepository) ReportUseCase {
	return &reportUseCase{
		dailyExpenditureRepo: dailyExpenditureRepo,
		transactionRepo:      transactionRepo,
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

func (uc *reportUseCase) GenerateReport(startDate time.Time, endDate time.Time) (*model.TransactionReport, error) {
	// total incomeTransaction
	income, err := uc.transactionRepo.SumIncomeTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}
	// total expendituresTransaction
	expenditureTransaction, err := uc.transactionRepo.SumOutcomeTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}
	// total expendituresDaily
	expenditureDaily, err := uc.dailyExpenditureRepo.GetTotalExpenditureByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	// get data transaction
	transaction, err := uc.transactionRepo.GetTransactionByRangeDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// total expendituresTransaction and expendituresDaily
	totalExpenditure := expenditureTransaction + expenditureDaily

	reportTransaction := &model.TransactionReport{
		StartDate:           startDate,
		EndDate:             endDate,
		TotalInTransaction:  income,
		TotalOutTransaction: totalExpenditure,
		Report:              []*model.TransactionReportDetail{},
	}

	for _, detTransaction := range transaction {
		detailReport := &model.TransactionReportDetail{
			InvoiceNumber:       detTransaction.InvoiceNumber,
			CustomerName:        detTransaction.Name,
			CompanyName:         detTransaction.Company,
			PhoneNumberCustomer: detTransaction.PhoneNumber,
			TxType:              detTransaction.TxType,
			Total:               detTransaction.Total,
		}
		reportTransaction.Report = append(reportTransaction.Report, detailReport)
	}

	return reportTransaction, nil
}

func (uc *reportUseCase) GenerateSalesReport(startDate time.Time, endDate time.Time) (*model.SalesReportOut, error) {
	tx_type := "out"

	transaction, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	total, err := uc.transactionRepo.SumIncomeTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}

	reportTransaction := &model.SalesReportOut{
		StartDate:  startDate,
		EndDate:    endDate,
		SalesTotal: total,
		Report:     []*model.TransactionReportDetail{},
	}

	for i, detTransaction := range transaction {
		detailReport := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       detTransaction.InvoiceNumber,
			CustomerName:        detTransaction.Name,
			Date:                detTransaction.Date,
			CompanyName:         detTransaction.Company,
			PhoneNumberCustomer: detTransaction.PhoneNumber,
			TxType:              detTransaction.TxType,
			Total:               detTransaction.Total,
		}
		reportTransaction.Report = append(reportTransaction.Report, detailReport)
	}

	return reportTransaction, nil
}

func (uc *reportUseCase) GenerateReceiptReport(startDate time.Time, endDate time.Time) (*model.ReceiptReport, error) {
	tx_type := "out"
	transaction, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	total, err := uc.transactionRepo.SumIncomeTransactionsWithType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}

	reportTransaction := &model.ReceiptReport{
		StartDate:    startDate,
		EndDate:      endDate,
		ReceiptTotal: total,
		Report:       []*model.TransactionReportDetail{},
	}
	for i, detTransaction := range transaction {
		debt := detTransaction.Total - detTransaction.PaymentAmount
		detailReport := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       detTransaction.InvoiceNumber,
			CustomerName:        detTransaction.Name,
			Date:                detTransaction.Date,
			CompanyName:         detTransaction.Company,
			PhoneNumberCustomer: detTransaction.PhoneNumber,
			TxType:              detTransaction.TxType,
			PaymentStatus:       detTransaction.PaymentStatus,
			Total:               detTransaction.Total,
			DebtTotal:           debt,
		}
		reportTransaction.Report = append(reportTransaction.Report, detailReport)
	}
	return reportTransaction, nil
}
