package usecase

import (
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/apperror"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"enigmacamp.com/final-project/team-4/track-prosto/utils/common"
)

type CustomerUseCase interface {
	CreateCustomer(*model.CustomerReqModel) error
	UpdateCustomer(*model.CustomerModel) error
	GetCustomerById(string) (*model.CustomerModel, error)
	GetAllCustomers() ([]*model.CustomerModel, error)
	DeleteCustomer(string) error
	GetAllCustomerTransactions(customer_id string) ([]*model.TransactionHeader, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
	companyRepo  repository.CompanyRepository
	txRepo repository.TransactionRepository
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository, companyRepo repository.CompanyRepository, txRepo repository.TransactionRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
		companyRepo:  companyRepo,
		txRepo: txRepo,
	}
}

func (cu *customerUseCase) CreateCustomer(customerReq *model.CustomerReqModel) error {
	now := time.Now()
	customerReq.CreatedAt = now
	customerReq.UpdatedAt = now
	var company model.Company
	customer := model.CustomerModel{
		Id:          customerReq.Id,
		FullName:    customerReq.FullName,
		Address:     customerReq.Address,
		PhoneNumber: customerReq.PhoneNumber,
		CreatedAt:   customerReq.CreatedAt,
		UpdatedAt:   customerReq.UpdatedAt,
		CreatedBy:   customerReq.CreatedBy,
		UpdatedBy:   customerReq.UpdatedBy,
	}
	if customerReq.Company.Address != "" || customerReq.Company.Email != "" || customerReq.Company.PhoneNumber != "" {
		existCompany, err := cu.companyRepo.GetCompanyByName(customerReq.Company.CompanyName)
		if err != nil {
			return fmt.Errorf("error on companyRepo.GetCompanyByName() : %w", err)
		}

		if existCompany == nil {
			company = model.Company{
				ID:          common.UuidGenerate(),
				CompanyName: customerReq.Company.CompanyName,
				Address:     customerReq.Company.Address,
				Email:       customerReq.Company.Email,
				PhoneNumber: customerReq.Company.PhoneNumber,
				IsActive:    customerReq.Company.IsActive,
				CreatedAt:   customerReq.CreatedAt,
				UpdatedAt:   customerReq.UpdatedAt,
				CreatedBy:   customerReq.CreatedBy,
				UpdatedBy:   customerReq.UpdatedBy,
			}
			err := cu.companyRepo.CreateCompany(&company)
			if err != nil {
				return fmt.Errorf("error on companyRepo.GetCompanyByName() : %w", err)
			}
			customer.CompanyId = company.ID
			return cu.customerRepo.CreateCustomer(&customer)
		}

		customer.CompanyId = existCompany.ID
		return cu.customerRepo.CreateCustomer(&customer)
	} else {
		existCompany, err := cu.companyRepo.GetCompanyByName(customerReq.Company.CompanyName)
		if err != nil {
			return fmt.Errorf("error on companyRepo.GetCompanyByName() : %w", err)
		}
		if existCompany == nil {
			return apperror.AppError{
				ErrorCode:    400,
				ErrorMassage: fmt.Sprintf("company with name : %v not available, please fill in the company data completely", customerReq.Company.CompanyName),
			}
		}
		customer.CompanyId = existCompany.ID
		return cu.customerRepo.CreateCustomer(&customer)
	}

}

func (cu *customerUseCase) UpdateCustomer(customer *model.CustomerModel) error {
	// Set updated_at timestamp
	customer.UpdatedAt = time.Now()
	existName, err := cu.customerRepo.GetCustomerByName(customer.FullName)
	if err != nil {
		return fmt.Errorf("error on customerRepo.GetCustomerByName() : %w", err)
	}
	if existName == nil {
		return apperror.AppError{
			ErrorCode:    404,
			ErrorMassage: fmt.Sprintf("Customer with name : %v not found", customer.FullName),
		}
	}

	return cu.customerRepo.UpdateCustomer(customer)
}

func (cu *customerUseCase) GetCustomerById(id string) (*model.CustomerModel, error) {
	customer, err := cu.customerRepo.GetCustomerById(id)
	if err != nil {
		return nil, fmt.Errorf("error oncustomerRepo.GetCustomerById() %w : ", err)

	}

	if customer == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data customer with id : %s not found", id),
		}
	}

	return customer, nil
}

func (cu *customerUseCase) GetAllCustomers() ([]*model.CustomerModel, error) {
	customers, err := cu.customerRepo.GetAllCustomer()
	if err != nil {
		return nil, fmt.Errorf("error oncustomerRepo.GetCustomerById() %w : ", err)

	}
	if customers == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data customer is empty",
		}
	}
	return customers, nil
}

func (cu *customerUseCase) DeleteCustomer(id string) error {
	existId, err := cu.customerRepo.GetCustomerById(id)
	if err != nil {
		return fmt.Errorf("error oncustomerRepo.GetCustomerById() %w : ", err)
	}
	if existId == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data customer with id : %s not found", id),
		}
	}
	return cu.customerRepo.DeleteCustomer(id)
}

func (cu *customerUseCase) GetAllCustomerTransactions(username string) ([]*model.TransactionHeader, error) {
	customerTransactions, err := cu.txRepo.GetAllTransactionsByCustomerUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error GetAllCustomerTransactions %w : ", err)

	}
	if customerTransactions == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data customer is empty",
		}
	}
	return customerTransactions, nil
}