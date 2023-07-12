package usecase

import (
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/apperror"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
}

type loginUsecaseImpl struct {
	userRepo repository.UserRepository
}

func (lu *loginUsecaseImpl) IsNameOrPassExist(name string, pass string) (*model.UserModel, error) {
	usr, err := lu.userRepo.GetUserByName(name)
	if err != nil {
		return nil, fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w", err)
	}
	if usr == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("user data with name : %s not found", name),
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass))
	if err != nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "password that you entered is incorrect.",
		}
	}
	return usr, nil

}

func NewLoginUsecase(userRepo repository.UserRepository) LoginUsecase {
	return &loginUsecaseImpl{
		userRepo: userRepo,
	}
}
