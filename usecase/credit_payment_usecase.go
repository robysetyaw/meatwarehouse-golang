package usecase

import (
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type CreditPaymentUseCase interface {
	CreateCreditPayment(payment *model.CreditPayment) error
	GetCreditPayments() ([]*model.CreditPayment, error)
	GetCreditPaymentByID(id string) (*model.CreditPayment, error)
	UpdateCreditPayment(payment *model.CreditPayment) error
}
type creditPaymentUseCase struct {
	creditPaymentRepo repository.CreditPaymentRepository
	transactionRepo repository.TransactionRepository
}

func NewCreditPaymentUseCase(creditPaymentRepo repository.CreditPaymentRepository, transactionRepo repository.TransactionRepository) CreditPaymentUseCase {
	return &creditPaymentUseCase{
		creditPaymentRepo: creditPaymentRepo,
		transactionRepo: transactionRepo,
	}
}

func (uc *creditPaymentUseCase) CreateCreditPayment(payment *model.CreditPayment) error {
	// Validasi atau logika bisnis sebelum membuat pembayaran kredit
	// ...
	transaction,err := uc.transactionRepo.GetByInvoiceNumber(payment.InvoiceNumber)
	if err != nil {
		return err
	}
	if transaction == nil {
		return fmt.Errorf("invoiceNumber Not Exist")
	}
	if transaction.PaymentStatus == "paid" {
		return fmt.Errorf("invoice Already Paid")
	}
	totalCredit, err := uc.creditPaymentRepo.GetTotalCredit(payment.InvoiceNumber)

	if err != nil {
		return  err
	}

	if totalCredit >= transaction.Total {
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

func (uc *creditPaymentUseCase) GetCreditPayments() ([]*model.CreditPayment, error) {
	payments, err := uc.creditPaymentRepo.GetAllCreditPayments()
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (uc *creditPaymentUseCase) GetCreditPaymentByID(id string) (*model.CreditPayment, error) {
	payment, err := uc.creditPaymentRepo.GetCreditPaymentByID(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (uc *creditPaymentUseCase) UpdateCreditPayment(payment *model.CreditPayment) error {
	// Validasi atau logika bisnis sebelum memperbarui pembayaran kredit
	// ...

	err := uc.creditPaymentRepo.UpdateCreditPayment(payment)
	if err != nil {
		return err
	}

	return nil
}

