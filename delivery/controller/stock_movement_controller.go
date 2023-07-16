package controller

import (
	"net/http"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
)

type StockMovementController struct {
	stockMovementUseCase usecase.StockMovementUseCase
}

func NewStockMovementController(r *gin.Engine, stockMovementUseCase usecase.StockMovementUseCase) *StockMovementController {
	controller := &StockMovementController{
		stockMovementUseCase: stockMovementUseCase,
	}
	r.GET("/report/stock-movement", controller.GetStockMovementReport)
	return controller
}

func (sm *StockMovementController) GetStockMovementReport(c *gin.Context)  {
	var request struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}
	stockMovementReports, err := sm.stockMovementUseCase.GenerateStockMovementReport(startDate,endDate)
	if err != nil {
		// Handle error
		return
	}

	c.JSON(http.StatusOK, stockMovementReports)
}
