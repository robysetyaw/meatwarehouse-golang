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
	
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
	}
}

func (uc *transactionUseCase) CreateTransaction(transaction *model.TransactionHeader) error {
	// Generate invoice number
	today := time.Now().Format("20060102")
	number, err := uc.transactionRepo.CountTransactions()
	if err != nil {
		return err
	}

	invoiceNumber := fmt.Sprintf("INV-%s-%04d", today, number )
	transaction.ID = common.UuidGenerate()
	transaction.IsActive = true
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	transaction.InvoiceNumber = invoiceNumber

	// Perform any business logic or validation before creating the transaction
	// ...

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
