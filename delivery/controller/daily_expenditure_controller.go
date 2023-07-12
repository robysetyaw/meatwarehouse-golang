package controller

import (
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/utils"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DailyExpenditureController struct {
	dailyExpenditureUseCase usecase.DailyExpenditureUseCase
}

func NewDailyExpenditureController(r *gin.Engine, deUC usecase.DailyExpenditureUseCase) *DailyExpenditureController {
	controller := &DailyExpenditureController{
		dailyExpenditureUseCase: deUC,
	}

	r.POST("/daily-expenditures",middleware.JWTAuthMiddleware(), controller.CreateDailyExpenditure)
	r.PUT("/daily-expenditures/:id",middleware.JWTAuthMiddleware(), controller.UpdateDailyExpenditure)
	r.GET("/daily-expenditures/:id",middleware.JWTAuthMiddleware(), controller.GetDailyExpenditureByID)
	r.GET("/daily-expenditures",middleware.JWTAuthMiddleware(), controller.GetAllDailyExpenditures)
	r.DELETE("/daily-expenditures/:id",middleware.JWTAuthMiddleware(), controller.DeleteDailyExpenditure)

	return controller
}

func (dec *DailyExpenditureController) CreateDailyExpenditure(c *gin.Context) {
	var expenditure model.DailyExpenditure
	if err := c.ShouldBindJSON(&expenditure); err != nil {
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

	userID := claims["user_id"].(string)
	userName := claims["username"].(string)
	expenditure.UserID = userID
	expenditure.CreatedBy = userName
	expenditure.ID = uuid.New().String()

	if err := dec.dailyExpenditureUseCase.CreateDailyExpenditure(&expenditure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create daily expenditure"})
		return
	}

	c.JSON(http.StatusOK, expenditure)
}

func (dec *DailyExpenditureController) UpdateDailyExpenditure(c *gin.Context) {
	expenditureID := c.Param("id")

	var expenditure model.DailyExpenditure
	if err := c.ShouldBindJSON(&expenditure); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenditure.ID = expenditureID

	if err := dec.dailyExpenditureUseCase.UpdateDailyExpenditure(&expenditure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update daily expenditure"})
		return
	}

	c.JSON(http.StatusOK, expenditure)
}

func (dec *DailyExpenditureController) GetDailyExpenditureByID(c *gin.Context) {
	expenditureID := c.Param("id")

	expenditure, err := dec.dailyExpenditureUseCase.GetDailyExpenditureByID(expenditureID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get daily expenditure"})
		return
	}

	if expenditure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daily expenditure not found"})
		return
	}

	c.JSON(http.StatusOK, expenditure)
}

func (dec *DailyExpenditureController) GetAllDailyExpenditures(c *gin.Context) {
	expenditures, err := dec.dailyExpenditureUseCase.GetAllDailyExpenditures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get daily expenditures"})
		return
	}

	c.JSON(http.StatusOK, expenditures)
}

func (dec *DailyExpenditureController) DeleteDailyExpenditure(c *gin.Context) {
	expenditureID := c.Param("id")

	if err := dec.dailyExpenditureUseCase.DeleteDailyExpenditure(expenditureID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete daily expenditure"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Daily expenditure deleted successfully"})
}
