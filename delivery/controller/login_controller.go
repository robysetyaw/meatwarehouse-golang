package controller

import (
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	// userUseCase  usecase.UserUseCase
	loginUseCase usecase.LoginUsecase
}

func NewLoginController(r *gin.Engine, loginUC usecase.LoginUsecase) {
	loginController := &LoginController{
		// userUseCase:  userUC,
		loginUseCase: loginUC,
	}

	r.POST("/login", loginController.Login)
}

func (uc *LoginController) Login(c *gin.Context) {
	var loginData model.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	if loginData.Username == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := uc.loginUseCase.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token ": token})
}
