package controller

import (
	"net/http"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase, router *gin.Engine) *UserController {
	controller := &UserController{
		userUseCase: userUseCase,
	}

	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.GET("/users", controller.GetAllUsers)

	return controller
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Mengambil data dari request
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mengenkripsi password sebelum menyimpannya
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user := &model.User{
		ID:        request.ID,
		Username:  request.Username,
		Password:  string(hashedPassword),
		IsActive:  request.IsActive,
		Role:      request.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: request.CreatedBy,
		UpdatedBy: request.UpdatedBy,
	}

	err = uc.userUseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// Mengambil ID pengguna dari URL parameter
	id := c.Param("id")

	// Mengambil data dari request
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := uc.userUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Mengenkripsi password baru jika ada
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	existingUser.Username = request.Username
	existingUser.IsActive = request.IsActive
	existingUser.Role = request.Role
	existingUser.UpdatedAt = time.Now()
	existingUser.UpdatedBy = request.UpdatedBy

	err = uc.userUseCase.UpdateUser(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// Mengambil ID pengguna dari URL parameter
	id := c.Param("id")

	err := uc.userUseCase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// Mengambil ID pengguna dari URL parameter
	id := c.Param("id")

	user, err := uc.userUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
