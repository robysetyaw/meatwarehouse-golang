package controller

import (
	"net/http"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUseCase usecase.ReportUseCase
}

func NewReportController(r *gin.Engine ,reportUseCase usecase.ReportUseCase) *ReportController {
	controller :=  &ReportController{
		reportUseCase: reportUseCase,
	}
	r.GET("/daily-expenditures/report",middleware.JWTAuthMiddleware(), controller.GenerateExpenditureReport)
	
	return controller
}

func (erc *ReportController) GenerateExpenditureReport(c *gin.Context) {
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
	
	report, err := erc.reportUseCase.GenerateExpenditureReport(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
