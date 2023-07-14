package controller

import (
	"errors"
	"fmt"
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/apperror"
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/utils"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"enigmacamp.com/final-project/team-4/track-prosto/utils/common"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerUsecase usecase.CustomerUseCase
}

func NewCustomerController(r *gin.Engine, customerUsecase usecase.CustomerUseCase) *CustomerController {
	controller := &CustomerController{
		customerUsecase: customerUsecase,
	}
	r.POST("/customer", middleware.JWTAuthMiddleware(), controller.CreateCustomer)
	r.GET("/customers", middleware.JWTAuthMiddleware(), controller.GetAllCustomer)
	r.GET("/customer/:id", middleware.JWTAuthMiddleware(), controller.GetCustomerByID)
	r.PUT("/customer/:id", middleware.JWTAuthMiddleware(), controller.UpdateCustomer)
	r.DELETE("/customer/:id", middleware.JWTAuthMiddleware(), controller.DeleteCustomer)
	return controller
}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	var customer model.CustomerReqModel
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.ExtractTokenFromAuthHeader(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		return
	}

	claims, err := utils.VerifyJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userName := claims["username"].(string)
	customer.CreatedBy = userName
	customer.Id = common.UuidGenerate()

	if err := cc.customerUsecase.CreateCustomer(&customer); err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("customerUsecase.CreateCustomer() : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"errorMessage": appError.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to create customer",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success insert data customer",
	})
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")

	var customer model.CustomerModel
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.ExtractTokenFromAuthHeader(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		return
	}

	claims, err := utils.VerifyJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	userName := claims["username"].(string)
	customer.UpdatedBy = userName
	customer.Id = customerID

	if err := cc.customerUsecase.UpdateCustomer(&customer); err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf(" cc.customerUsecase.UpdateCustomer() : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.InsertService() 2 : %v ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to update customer",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update data company",
	})
}

func (cc *CustomerController) GetAllCustomer(c *gin.Context) {
	customers, err := cc.customerUsecase.GetAllCustomers()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf(" cc.customerUsecase.GetAllCustomers() : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.InsertService() 2 : %v ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to get customers",
			})
		}
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
	customerId := c.Param("id")

	expenditure, err := cc.customerUsecase.GetCustomerById(customerId)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf(" cc.customerUsecase.GetCustomerById() : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"errorMessage": appError.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to get customer",
			})
		}
		return
	}
	c.JSON(http.StatusOK, expenditure)
}

func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
	customerId := c.Param("id")

	if err := cc.customerUsecase.DeleteCustomer(customerId); err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf(" cc.customerUsecase.DeleteCustomer() : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"errorMessage": appError.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to delete customer",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}
