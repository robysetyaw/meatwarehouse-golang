package controller

import (
	"errors"
	"fmt"
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/apperror"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
	"enigmacamp.com/final-project/team-4/track-prosto/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(srv *gin.Engine, userUseCase usecase.UserUseCase) *UserController {
	controller := &UserController{
		userUseCase: userUseCase,
	}

	srv.POST("/users", controller.CreateUser)
	// srv.PUT("/users/:id", controller.UpdateUser)
	// srv.DELETE("/users/:id", controller.DeleteUser)
	// srv.GET("/users/:id", controller.GetUserByID)
	// srv.GET("/users", controller.GetAllUsers)

	return controller
}

func (uc *UserController) CreateUser(c *gin.Context) {
	usr := &model.UserModel{}
	usr.ID = utils.UuidGenerate()
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = uc.userUseCase.CreateUser(usr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("usrHandler.usrUsecase() 1 : %v ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("usrHandler.usrUsecase() 2 : %v ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// func (uc *UserController) UpdateUser(c *gin.Context) {
// 	// Mengambil ID pengguna dari URL parameter
// 	id := c.Param("id")

// 	// Mengambil data dari request
// 	var request model.UserModel
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	existingUser, err := uc.userUseCase.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}
// 	if existingUser == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Mengenkripsi password baru jika ada
// 	if request.Password != "" {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 			return
// 		}
// 		existingUser.Password = string(hashedPassword)
// 	}

// 	existingUser.Username = request.Username
// 	existingUser.IsActive = request.IsActive
// 	existingUser.Role = request.Role
// 	existingUser.UpdatedAt = time.Now()
// 	existingUser.UpdatedBy = request.UpdatedBy

// 	err = uc.userUseCase.UpdateUser(existingUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
// }

// func (uc *UserController) DeleteUser(c *gin.Context) {
// 	// Mengambil ID pengguna dari URL parameter
// 	id := c.Param("id")

// 	err := uc.userUseCase.DeleteUser(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// Mengambil ID pengguna dari URL parameter
// 	id := c.Param("id")

// 	user, err := uc.userUseCase.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
// 		return
// 	}
// 	if user == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// func (uc *UserController) GetAllUsers(c *gin.Context) {
// 	users, err := uc.userUseCase.GetAllUsers()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }
