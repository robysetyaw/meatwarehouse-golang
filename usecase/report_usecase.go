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
	GenerateDebtAccountsPayableReport(startDate time.Time, endDate time.Time) (*model.DebtAccountsPayableReport, error)
	GenerateProfitLossStatement(startDate time.Time, endDate time.Time) (*model.ProfitAndLossStatement, error)
	GenerateCashFlowStatement(startDate time.Time, endDate time.Time) (*model.CashFlowStatement, error)
	GenerateConsolidatedReport(startDate time.Time, endDate time.Time) (*model.ConsolidatedReport, error)
}

type reportUseCase struct {
	transactionRepo      repository.TransactionRepository
	dailyExpenditureRepo repository.DailyExpenditureRepository
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

	for i, expenditure := range expenditures {
		dailyReport := &model.DailyExpenditureReport{
			No:          i+1,
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
	total, err := uc.transactionRepo.SumPaymentTransactionsWithType(startDate, endDate, tx_type)
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
			PaymentAmount:       detTransaction.PaymentAmount,
			PaymentStatus:       detTransaction.PaymentStatus,
			Total:               detTransaction.Total,
			DebtTotal:           debt,
		}
		reportTransaction.Report = append(reportTransaction.Report, detailReport)
	}
	return reportTransaction, nil
}

func (uc *reportUseCase) GenerateDebtAccountsPayableReport(startDate time.Time, endDate time.Time) (*model.DebtAccountsPayableReport, error) {
	tx_type := "out"
	status := "unpaid"
	transactionIn, err := uc.transactionRepo.GetTransactionByRangeDateWithTxTypeAndPaid(startDate, endDate, tx_type, status)
	if err != nil {
		return nil, err
	}

	payIn, err := uc.transactionRepo.SumPaymentTransactionsWithTypeAndStatus(startDate, endDate, tx_type,status)
	if err != nil {
		return nil, err
	}
	totalIn,err := uc.transactionRepo.SumTotalTransactionWithTypeAndStatus(startDate,endDate,tx_type,status)
	if err != nil {
		return nil, err
	}
	tx_type = "in"
	transactionOut, err := uc.transactionRepo.GetTransactionByRangeDateWithTxTypeAndPaid(startDate, endDate, tx_type, status)
	if err != nil {
		return nil, err
	}
	payOut, err := uc.transactionRepo.SumPaymentTransactionsWithTypeAndStatus(startDate, endDate, tx_type, status)
	if err != nil {
		return nil, err
	}
	totalOut,err := uc.transactionRepo.SumTotalTransactionWithTypeAndStatus(startDate,endDate,tx_type, status)
	if err != nil {
		return nil, err
	}
	reportTransaction := &model.DebtAccountsPayableReport{
		StartDate:        startDate,
		EndDate:          endDate,
		ReceivablesTotal: totalIn - payIn ,
		DebtTotal:        totalOut - payOut,
		Receivables:      []*model.DebtAccountsPayableReportDetail{},
		Debt:             []*model.DebtAccountsPayableReportDetail{},
	}
	for i, detTransaction := range transactionIn {
		debt := detTransaction.Total - detTransaction.PaymentAmount
		detailReport := &model.DebtAccountsPayableReportDetail{
			No:                  i + 1,
			InvoiceNumber:       detTransaction.InvoiceNumber,
			CustomerName:        detTransaction.Name,
			Date:                detTransaction.Date,
			CompanyName:         detTransaction.Company,
			PhoneNumberCustomer: detTransaction.PhoneNumber,
			TxType:              detTransaction.TxType,
			PaymentStatus:       detTransaction.PaymentStatus,
			Debt:                debt,
		}
		reportTransaction.Receivables = append(reportTransaction.Receivables, detailReport)
	}
	for i, detTransaction := range transactionOut {
		debt := detTransaction.Total - detTransaction.PaymentAmount
		detailReport := &model.DebtAccountsPayableReportDetail{
			No:                  i + 1,
			InvoiceNumber:       detTransaction.InvoiceNumber,
			CustomerName:        detTransaction.Name,
			Date:                detTransaction.Date,
			CompanyName:         detTransaction.Company,
			PhoneNumberCustomer: detTransaction.PhoneNumber,
			TxType:              detTransaction.TxType,
			PaymentStatus:       detTransaction.PaymentStatus,
			Debt:                debt,
		}
		reportTransaction.Debt = append(reportTransaction.Debt, detailReport)
	}
	return reportTransaction, nil
}

func (uc *reportUseCase) GenerateProfitLossStatement(startDate time.Time, endDate time.Time) (*model.ProfitAndLossStatement, error) {

	tx_type := "out"

	sales, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	totalSales, err := uc.transactionRepo.SumIncomeTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}
	tx_type = "in"
	transactionOut, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	totalOut, err := uc.transactionRepo.SumPaymentTransactionsWithType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}

	totalExpenditure, err := uc.dailyExpenditureRepo.GetTotalExpenditureByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	expenditures, err := uc.dailyExpenditureRepo.GetExpendituresByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	PnL := &model.ProfitAndLossStatement{
		StartDate:          startDate,
		EndDate:            endDate,
		SalesTotal:         totalSales,
		TotalExpenditure:   totalExpenditure,
		TotalInTransaction: totalOut,
		Profit:             totalSales - (totalExpenditure + totalOut),
		Sales:              []*model.TransactionReportDetail{},
		TransactionIn:      []*model.TransactionReportDetail{},
		Expenditures:       []*model.DailyExpenditureReport{},
	}

	for i, s := range sales {
		detailReport := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       s.InvoiceNumber,
			Date:                s.Date,
			CustomerName:        s.Name,
			CompanyName:         s.Company,
			PhoneNumberCustomer: s.PhoneNumber,
			TxType:              s.TxType,
			Total:               s.Total,
		}
		PnL.Sales = append(PnL.Sales, detailReport)
	}
	for i, ti := range transactionOut {
		detailTransactionIn := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       ti.InvoiceNumber,
			Date:                ti.Date,
			CustomerName:        ti.Name,
			CompanyName:         ti.Company,
			PhoneNumberCustomer: ti.PhoneNumber,
			TxType:              ti.TxType,
			Total:               ti.Total,
		}
		PnL.TransactionIn = append(PnL.TransactionIn, detailTransactionIn)
	}
	for i, expenditure := range expenditures {
		dailyReport := &model.DailyExpenditureReport{
			No:          i,
			ID:          expenditure.ID,
			UserID:      expenditure.UserID,
			Username:    expenditure.Username,
			Amount:      expenditure.Amount,
			Description: expenditure.Description,
			CreatedAt:   expenditure.CreatedAt,
			UpdatedAt:   expenditure.UpdatedAt,
			Date:        expenditure.Date,
		}

		PnL.Expenditures = append(PnL.Expenditures, dailyReport)
	}
	return PnL, nil
}

func (uc *reportUseCase) GenerateCashFlowStatement(startDate time.Time, endDate time.Time) (*model.CashFlowStatement, error) {
	tx_type := "in"
	totalCashOut, err := uc.transactionRepo.SumPaymentTransactionsWithType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	paymentIn, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	tx_type = "out"
	totalCashIn, err := uc.transactionRepo.SumPaymentTransactionsWithType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}
	paymentOut, err := uc.transactionRepo.GetTransactionByRangeDateWithTxType(startDate, endDate, tx_type)
	if err != nil {
		return nil, err
	}

	totalExpenditure, err := uc.dailyExpenditureRepo.GetTotalExpenditureByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	expenditures, err := uc.dailyExpenditureRepo.GetExpendituresByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	reportTransaction := &model.CashFlowStatement{
		StartDate:        startDate,
		EndDate:          endDate,
		TotalPaymentIn:   totalCashIn,
		TotalPaymentOut:  totalCashOut,
		TotalExpenditure: totalExpenditure,
		Cash:             totalCashIn - (totalExpenditure + totalCashOut),
		PaymentIn:        []*model.TransactionReportDetail{},
		PaymentOut:       []*model.TransactionReportDetail{},
		Expenditure:      []*model.DailyExpenditureReport{},
	}

	for i, s := range paymentIn {
		detailReport := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       s.InvoiceNumber,
			Date:                s.Date,
			CustomerName:        s.Name,
			CompanyName:         s.Company,
			PhoneNumberCustomer: s.PhoneNumber,
			TxType:              s.TxType,
			Total:               s.Total,
		}
		reportTransaction.PaymentIn = append(reportTransaction.PaymentIn, detailReport)
	}
	for i, ti := range paymentOut {
		detailTransactionIn := &model.TransactionReportDetail{
			No:                  i + 1,
			InvoiceNumber:       ti.InvoiceNumber,
			Date:                ti.Date,
			CustomerName:        ti.Name,
			CompanyName:         ti.Company,
			PhoneNumberCustomer: ti.PhoneNumber,
			TxType:              ti.TxType,
			Total:               ti.Total,
		}
		reportTransaction.PaymentOut = append(reportTransaction.PaymentOut, detailTransactionIn)
	}
	for i, expenditure := range expenditures {
		dailyReport := &model.DailyExpenditureReport{
			No:          i,
			ID:          expenditure.ID,
			UserID:      expenditure.UserID,
			Username:    expenditure.Username,
			Amount:      expenditure.Amount,
			Description: expenditure.Description,
			CreatedAt:   expenditure.CreatedAt,
			UpdatedAt:   expenditure.UpdatedAt,
			Date:        expenditure.Date,
		}

		reportTransaction.Expenditure = append(reportTransaction.Expenditure, dailyReport)
	}
	return reportTransaction, nil
}

func (uc *reportUseCase) GenerateConsolidatedReport(startDate time.Time, endDate time.Time) (*model.ConsolidatedReport, error) {
	expenditureReport, err := uc.GenerateExpenditureReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	transactionReport, err := uc.GenerateReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	salesReport, err := uc.GenerateSalesReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	receiptReport, err := uc.GenerateReceiptReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	debtAccountsPayableReport, err := uc.GenerateDebtAccountsPayableReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	profitLossStatement, err := uc.GenerateProfitLossStatement(startDate, endDate)
	if err != nil {
		return nil, err
	}

	cashFlowStatement, err := uc.GenerateCashFlowStatement(startDate, endDate)
	if err != nil {
		return nil, err
	}

	consolidatedReport := &model.ConsolidatedReport{
		ExpenditureReport:         expenditureReport,
		TransactionReport:         transactionReport,
		SalesReport:               salesReport,
		ReceiptReport:             receiptReport,
		DebtAccountsPayableReport: debtAccountsPayableReport,
		ProfitLossStatement:       profitLossStatement,
		CashFlowStatement:         cashFlowStatement,
	}

	return consolidatedReport, nil
}
