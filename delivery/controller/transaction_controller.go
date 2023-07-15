package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionController(r *gin.Engine, transactionUseCase usecase.TransactionUseCase) *TransactionController {
	controller := &TransactionController{
		transactionUseCase: transactionUseCase,
	}

	r.POST("/transactions", middleware.JWTAuthMiddleware(), controller.CreateTransaction)
	r.GET("/transactions/:id", middleware.JWTAuthMiddleware(), controller.GetTransactionByID)
	r.GET("/transactions", middleware.JWTAuthMiddleware(), controller.GetAllTransactions)
	r.DELETE("/transactions/:id", middleware.JWTAuthMiddleware(), controller.DeleteTransaction)

	return controller
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var request model.TransactionHeader
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.transactionUseCase.CreateTransaction(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

func (tc *TransactionController) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	transaction, err := tc.transactionUseCase.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetAllTransactions(c *gin.Context) {
	transactions, err := tc.transactionUseCase.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	err := tc.transactionUseCase.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
