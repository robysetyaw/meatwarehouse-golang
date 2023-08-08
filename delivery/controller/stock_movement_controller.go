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

func (sm *StockMovementController) GetStockMovementReport(c *gin.Context) {
	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}
	
	stockMovementReports, err := sm.stockMovementUseCase.GenerateStockMovementReport(startDate, endDate)
	if err != nil {
		// Handle error
		return
	}

	c.JSON(http.StatusOK, stockMovementReports)
}
