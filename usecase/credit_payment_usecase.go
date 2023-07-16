package usecase

import (
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type CreditPaymentUseCase struct {
	creditPaymentRepo repository.CreditPaymentRepository
	transactionRepo repository.TransactionRepository
}

func NewCreditPaymentUseCase(creditPaymentRepo repository.CreditPaymentRepository, transactionRepo repository.TransactionRepository) *CreditPaymentUseCase {
	return &CreditPaymentUseCase{
		creditPaymentRepo: creditPaymentRepo,
		transactionRepo: transactionRepo,

	}
}

func (uc *CreditPaymentUseCase) CreateCreditPayment(payment *model.CreditPayment) error {
	// Validasi atau logika bisnis sebelum membuat pembayaran kredit
	// ...
	transaction,err := uc.transactionRepo.GetByInvoiceNumber(payment.InvoiceNumber)
	if err != nil {
		return err
	}
	if transaction == nil {
		return fmt.Errorf("InvoiceNumber Not Exist")
	}
	if transaction.PaymentStatus == "paid" {
		return fmt.Errorf("Invoice Already Paid")
	}
	totalCredit, err := uc.creditPaymentRepo.GetTotalCredit(payment.InvoiceNumber)

	if err != nil {
		return  err
	}

	if (totalCredit+transaction.PaymentAmount) >= transaction.Total {
		err = uc.transactionRepo.UpdateStatusInvoicePaid(transaction.ID)
		if err != nil {
			return err
		}
	}

	err = uc.creditPaymentRepo.CreateCreditPayment(payment)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CreditPaymentUseCase) GetCreditPayments() ([]*model.CreditPayment, error) {
	payments, err := uc.creditPaymentRepo.GetAllCreditPayments()
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (uc *CreditPaymentUseCase) GetCreditPaymentByID(id string) (*model.CreditPayment, error) {
	payment, err := uc.creditPaymentRepo.GetCreditPaymentByID(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (uc *CreditPaymentUseCase) UpdateCreditPayment(payment *model.CreditPayment) error {
	// Validasi atau logika bisnis sebelum memperbarui pembayaran kredit
	// ...

	err := uc.creditPaymentRepo.UpdateCreditPayment(payment)
	if err != nil {
		return err
	}

	return nil
}

