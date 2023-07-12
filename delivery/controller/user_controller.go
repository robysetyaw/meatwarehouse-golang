package controller

import (
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/middleware"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(r *gin.Engine, userUC usecase.UserUseCase) {
	userController := &UserController{
		userUseCase: userUC,
	}

	r.POST("/users", middleware.JWTAuthMiddleware(),userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	// r.GET("/users/:id", userController.GetUserByID)
	r.GET("/users/:username", userController.GetUserByUsername)
	r.GET("/users",middleware.JWTAuthMiddleware(), userController.GetAllUsers)
	r.DELETE("/users/:id", userController.DeleteUser)
}
func (uc *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := uc.userUseCase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = userID

	if err := uc.userUseCase.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := uc.userUseCase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := uc.userUseCase.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	if err := uc.userUseCase.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

