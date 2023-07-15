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
	// Validate transaction data
	if err := validateTransaction(transaction); err != nil {
		return err
	}

	// Set created and updated timestamps
	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	// Set created by and updated by
	// Replace with appropriate logic to get the user ID or username
	transaction.CreatedBy = "user123"
	transaction.UpdatedBy = "user123"

	// Generate unique ID for the transaction
	transaction.ID = common.UuidGenerate()

	// Create the transaction
	if err := uc.transactionRepo.CreateTransaction(transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return nil
}

func validateTransaction(transaction *model.TransactionHeader) error {
	// Add your validation logic here
	// Return an error if the transaction data is not valid
	return nil
}

func generateTransactionID() string {
	// Replace with your custom logic to generate a unique transaction ID
	// Example: return "TX-202307120001"
	return "TX-202307120001"
}
