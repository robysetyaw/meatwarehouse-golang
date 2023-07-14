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
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
	companyRepo  repository.CompanyRepository
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository, companyRepo repository.CompanyRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
		companyRepo:  companyRepo,
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
	existName, err := cu.customerRepo.

	return cu.customerRepo.UpdateCustomer(customer)
}

func (cu *customerUseCase) GetCustomerById(id string) (*model.CustomerModel, error) {
	return cu.customerRepo.GetCustomerById(id)
}

func (cu *customerUseCase) GetAllCustomers() ([]*model.CustomerModel, error) {
	return cu.customerRepo.GetAllCustomer()
}

func (cu *customerUseCase) DeleteCustomer(id string) error {
	return cu.customerRepo.DeleteCustomer(id)
}
