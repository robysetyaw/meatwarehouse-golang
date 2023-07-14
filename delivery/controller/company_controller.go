package controller

import (
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/utils"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"enigmacamp.com/final-project/team-4/track-prosto/utils/common"
	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyUseCase usecase.CompanyUseCase
}

func NewCompanyController(r *gin.Engine, companyUseCase usecase.CompanyUseCase) *CompanyController {
	controller := &CompanyController{
		companyUseCase: companyUseCase,
	}
	r.POST("/company", middleware.JWTAuthMiddleware(), controller.CreateCompany)
	r.PUT("/company", middleware.JWTAuthMiddleware(), controller.UpdateCompany)
	r.GET("/company/:id", middleware.JWTAuthMiddleware(), controller.GetCompanyById)
	r.GET("/companies", middleware.JWTAuthMiddleware(), controller.GetAllCompany)
	r.DELETE("/company/:id", middleware.JWTAuthMiddleware(), controller.DeleteCompany)

	return controller
}

func (cc *CompanyController) CreateCompany(c *gin.Context) {
	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
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
	company.CreatedBy = userName
	company.ID = common.UuidGenerate()
	company.IsActive = true

	if err := cc.companyUseCase.CreateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success insert data company",
	})
}

func (cc *CompanyController) UpdateCompany(c *gin.Context) {
	companyID := c.Param("id")

	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
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
	company.UpdatedBy = userName
	company.ID = companyID

	if err := cc.companyUseCase.UpdateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update update data company",
	})
}

func (cc *CompanyController) GetCompanyById(c *gin.Context) {
	companyId := c.Param("id")

	company, err := cc.companyUseCase.GetCompanyById(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get company"})
		return
	}

	if company == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (cc *CompanyController) GetAllCompany(c *gin.Context) {
	companies, err := cc.companyUseCase.GetAllCompany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (cc *CompanyController) DeleteCompany(c *gin.Context) {
	companyId := c.Param("id")

	if err := cc.companyUseCase.DeleteCompany(companyId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
