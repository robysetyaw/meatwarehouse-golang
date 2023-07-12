package usecase

import (
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/apperror"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(*model.UserModel) error
}
type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{
		userRepository: userRepository,
	}
}

func (uc *userUseCaseImpl) CreateUser(user *model.UserModel) error {
	isNameExist, err := uc.userRepository.GetUserByName(user.Username)
	if err != nil {
		return fmt.Errorf("userUseCaseImplImpl.InsertService() : %w", err)
	}

	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data with username : %v is exist", user.Username),
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf(" bcrypt.GenerateFromPassword() : %w", err)
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	return uc.userRepository.CreateUser(user)
}

// func (uc *userUseCaseImpl) UpdateUser(user *model.UserModel) error {
// 	// Lakukan validasi atau logika bisnis lainnya sebelum mengupdate pengguna di dalam repository
// 	// ...

// 	err := uc.userRepository.UpdateUser(user)
// 	if err != nil {
// 		// Tangani kesalahan jika terjadi kesalahan saat mengupdate pengguna
// 		// ...
// 		return err
// 	}

// 	return nil
// }

// func (uc *userUseCaseImpl) DeleteUser(userID string) error {
// 	// Lakukan validasi atau logika bisnis lainnya sebelum menghapus pengguna di dalam repository
// 	// ...

// 	err := uc.userRepository.DeleteUser(userID)
// 	if err != nil {
// 		// Tangani kesalahan jika terjadi kesalahan saat menghapus pengguna
// 		// ...
// 		return err
// 	}

// 	return nil
// }

// func (uc *userUseCaseImpl) GetUserByID(userID string) (*model.UserModel, error) {
// 	user, err := uc.userRepository.GetUserByID(userID)
// 	if err != nil {
// 		// Tangani kesalahan jika terjadi kesalahan saat mengambil pengguna berdasarkan ID
// 		// ...
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (uc *userUseCaseImpl) GetAllUsers() ([]*model.UserModel, error) {
// 	users, err := uc.userRepository.GetAllUsers()
// 	if err != nil {
// 		// Tangani kesalahan jika terjadi kesalahan saat mengambil daftar pengguna
// 		// ...
// 		return nil, err
// 	}

// 	return users, nil
// }
