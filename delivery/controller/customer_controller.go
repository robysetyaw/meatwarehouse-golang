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
			fmt.Printf("ServiceHandler.InsertService() 1 : %v ", err.Error())
			c.JSON(http.StatusBadGateway, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.InsertService() 2 : %v ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "failed to create customer",
			})
		}
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) GetAllCustomer(c *gin.Context) {
	customers, err := cc.customerUsecase.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get daily companies"})
		return
	}

	c.JSON(http.StatusOK, customers)
}
