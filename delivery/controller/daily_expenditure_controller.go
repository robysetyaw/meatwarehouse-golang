package controller

import (
	"net/http"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
)

type DailyExpenditureController struct {
	dailyExpenditureUseCase *usecase.DailyExpenditureUseCase
}



func NewDailyExpenditureController(dailyExpenditureUseCase *usecase.DailyExpenditureUseCase, router *gin.Engine) *DailyExpenditureController {
	controller := &DailyExpenditureController{
		dailyExpenditureUseCase: dailyExpenditureUseCase,
	}

	router.POST("/daily-expenditures", controller.CreateDailyExpenditure)
	router.PUT("/daily-expenditures/:id", controller.UpdateDailyExpenditure)
	router.DELETE("/daily-expenditures/:id", controller.DeleteDailyExpenditure)
	router.GET("/daily-expenditures/:id", controller.GetDailyExpenditureByID)
	router.GET("/daily-expenditures", controller.GetAllDailyExpenditures)

	return controller
}

func (dc *DailyExpenditureController) CreateDailyExpenditure(c *gin.Context) {
	// Mengambil data dari request
	var request model.DailyExpenditure
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dailyExpenditure := &model.DailyExpenditure{
		ID:          request.ID,
		UserID:      request.UserID,
		Amount:      request.Amount,
		Description: request.Description,
		IsActive:    request.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   request.CreatedBy,
		UpdatedBy:   request.UpdatedBy,
	}

	err := dc.dailyExpenditureUseCase.CreateDailyExpenditure(dailyExpenditure)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create daily expenditure"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Daily expenditure created successfully"})
}

func (dc *DailyExpenditureController) UpdateDailyExpenditure(c *gin.Context) {
	// Mengambil ID pengeluaran harian dari URL parameter
	id := c.Param("id")

	// Mengambil data dari request
	var request model.DailyExpenditure
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingDailyExpenditure, err := dc.dailyExpenditureUseCase.GetDailyExpenditureByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update daily expenditure"})
		return
	}
	if existingDailyExpenditure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daily expenditure not found"})
		return
	}

	existingDailyExpenditure.UserID = request.UserID
	existingDailyExpenditure.Amount = request.Amount
	existingDailyExpenditure.Description = request.Description
	existingDailyExpenditure.IsActive = request.IsActive
	existingDailyExpenditure.UpdatedAt = time.Now()
	existingDailyExpenditure.UpdatedBy = request.UpdatedBy

	err = dc.dailyExpenditureUseCase.UpdateDailyExpenditure(existingDailyExpenditure)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update daily expenditure"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Daily expenditure updated successfully"})
}

func (dc *DailyExpenditureController) DeleteDailyExpenditure(c *gin.Context) {
	// Mengambil ID pengeluaran harian dari URL parameter
	id := c.Param("id")

	err := dc.dailyExpenditureUseCase.DeleteDailyExpenditure(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete daily expenditure"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Daily expenditure deleted successfully"})
}

func (dc *DailyExpenditureController) GetDailyExpenditureByID(c *gin.Context) {
	// Mengambil ID pengeluaran harian dari URL parameter
	id := c.Param("id")

	dailyExpenditure, err := dc.dailyExpenditureUseCase.GetDailyExpenditureByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve daily expenditure"})
		return
	}
	if dailyExpenditure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daily expenditure not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dailyExpenditure})
}

func (dc *DailyExpenditureController) GetAllDailyExpenditures(c *gin.Context) {
	dailyExpenditures, err := dc.dailyExpenditureUseCase.GetAllDailyExpenditures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve daily expenditures"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dailyExpenditures})
}
