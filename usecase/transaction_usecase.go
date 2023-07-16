package usecase

import (
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"enigmacamp.com/final-project/team-4/track-prosto/utils/common"
)

type TransactionUseCase interface {
	CreateTransaction(transaction *model.TransactionHeader) error
	GetAllTransactions() ([]*model.TransactionHeader, error)
	GetTransactionByID(id string) (*model.TransactionHeader, error)
	DeleteTransaction(id string) error
	GetTransactionByInvoiceNumber(inv_number string) (*model.TransactionHeader, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	customerRepo    repository.CustomerRepository
	meatRepo        repository.MeatRepository
	companyRepo     repository.CompanyRepository
	creditPaymentRepo   repository.CreditPaymentRepository
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepository, customerRepo repository.CustomerRepository, meatRepo repository.MeatRepository, companyRepo repository.CompanyRepository, creditPaymentRepo   repository.CreditPaymentRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
		customerRepo:    customerRepo,
		meatRepo:        meatRepo,
		companyRepo:     companyRepo,
		creditPaymentRepo: creditPaymentRepo,
	}
}

func (uc *transactionUseCase) CreateTransaction(transaction *model.TransactionHeader) error {
	// Generate invoice number
	today := time.Now().Format("20060102")
	todayDate := time.Now().Format("2006-01-02")
	number, err := uc.transactionRepo.CountTransactions()
	if err != nil {
		return err
	}

	customer, err := uc.customerRepo.GetCustomerByName(transaction.Name)
	if err != nil {
		return fmt.Errorf("failed to get customer by name: %w", err)
	}

	company, err := uc.companyRepo.GetCompanyById(customer.CompanyId)
	if err != nil {
		return fmt.Errorf("failed to get customer by name: %w", err)
	}

	invoiceNumber := fmt.Sprintf("INV-%s-%04d", today, number)
	transaction.ID = common.UuidGenerate()
	transaction.Date = todayDate
	transaction.IsActive = true
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	transaction.InvoiceNumber = invoiceNumber
	transaction.CustomerID = customer.Id
	transaction.Address = customer.Address
	transaction.PhoneNumber = customer.PhoneNumber
	transaction.Company = company.CompanyName
	transaction.CreatedBy = "admin"
	transaction.UpdatedBy = "admin"
	transaction.PaymentStatus = "paid"
	// Perform any business logic or validation before creating the transaction
	// ...

	for _, detail := range transaction.TransactionDetails {
		meat, err := uc.meatRepo.GetMeatByName(detail.MeatName)
		if err != nil {
			return err
		}
		if meat == nil {
			return fmt.Errorf("meat name %s not found", meat.Name)
		}
		detail.ID = common.UuidGenerate()
		detail.MeatID = meat.ID
		detail.TransactionID = transaction.ID

		if detail.Qty >= meat.Stock {
			return fmt.Errorf("insufficient stock for %s", detail.MeatName)
		}

		if transaction.TxType == "in" {
			err := uc.meatRepo.IncreaseStock(meat.ID, detail.Qty)
			if err != nil {
				return err
			}
		}
		if transaction.TxType == "out" {
			err = uc.meatRepo.ReduceStock(meat.ID, detail.Qty)
			if err != nil {
				return err
			}
		}

	}
	transaction.CalulatedTotal()
	newTotal := uc.UpdateTotalTransaction(transaction)
	
	if transaction.PaymentAmount>newTotal {
		return fmt.Errorf("amount is large than total transaction")
	}

	if newTotal > transaction.PaymentAmount {
		transaction.PaymentStatus = "unpaid"
	}
	
	uc.creditPaymentRepo.CreateCreditPayment(&model.CreditPayment{
		ID: common.UuidGenerate(),
		InvoiceNumber: transaction.InvoiceNumber,
		Amount: transaction.PaymentAmount,
		PaymentDate: transaction.Date,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.CreatedAt,
		CreatedBy: transaction.CreatedBy,
		UpdatedBy: transaction.CreatedBy,
	})
	
	// Create transaction header
	if err := uc.transactionRepo.CreateTransactionHeader(transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

func (uc *transactionUseCase) GetAllTransactions() ([]*model.TransactionHeader, error) {
	return uc.transactionRepo.GetAllTransactions()
}

func (uc *transactionUseCase) GetTransactionByID(id string) (*model.TransactionHeader, error) {
	transaction, err := uc.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return transaction, nil
}

func (uc *transactionUseCase) DeleteTransaction(id string) error {
	return uc.transactionRepo.DeleteTransaction(id)
}

func (uc *transactionUseCase) UpdateTotalTransaction(transaction *model.TransactionHeader) float64 {
	var newTotal float64
	for _, detail := range transaction.TransactionDetails {
		detail.Total = detail.Price * detail.Qty
		newTotal = newTotal + detail.Total
	}

	return newTotal
}

func (uc *transactionUseCase) GetTransactionByInvoiceNumber(inv_number string) (*model.TransactionHeader, error) {
	transaction, err := uc.transactionRepo.GetByInvoiceNumber(inv_number)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return transaction, nil
}
