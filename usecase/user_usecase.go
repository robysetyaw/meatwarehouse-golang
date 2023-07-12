package usecase

import (
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type UserUseCase interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepo,
	}
}

func (uc *userUseCase) CreateUser(user *model.User) error {
	// Implement any business logic or validation before creating the user
	// You can also perform data manipulation or enrichment if needed

	err := uc.userRepository.CreateUser(user)
	if err != nil {
		// Handle any repository errors or perform error logging
		return err
	}

	return nil
}

func (uc *userUseCase) UpdateUser(user *model.User) error {
	// Implement any business logic or validation before updating the user
	// You can also perform data manipulation or enrichment if needed

	err := uc.userRepository.UpdateUser(user)
	if err != nil {
		// Handle any repository errors or perform error logging
		return err
	}

	return nil
}

func (uc *userUseCase) GetUserByID(id string) (*model.User, error) {
	user, err := uc.userRepository.GetUserByID(id)
	if err != nil {
		// Handle any repository errors or perform error logging
		return nil, err
	}

	// Perform any additional data processing or transformation if needed

	return user, nil
}

func (uc *userUseCase) GetAllUsers() ([]*model.User, error) {
	users, err := uc.userRepository.GetAllUsers()
	if err != nil {
		// Handle any repository errors or perform error logging
		return nil, err
	}

	// Perform any additional data processing or transformation if needed

	return users, nil
}

func (uc *userUseCase) DeleteUser(id string) error {
	// Implement any business logic or validation before deleting the user

	err := uc.userRepository.DeleteUser(id)
	if err != nil {
		// Handle any repository errors or perform error logging
		return err
	}

	return nil
}
