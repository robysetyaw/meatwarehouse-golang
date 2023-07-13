package usecase

import (
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"enigmacamp.com/final-project/team-4/track-prosto/utils/authutil"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
	Login(string, string) (string, error)
}

type loginUsecaseImpl struct {
	userRepo repository.UserRepository
}

func (uc *loginUsecaseImpl) Login(username, password string) (string, error) {
	// Mengecek apakah pengguna dengan username tersebut ada di penyimpanan data
	user, err := uc.userRepo.GetUserByName(username)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Verifikasi password pengguna dengan menggunakan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	// Menghasilkan token JWT
	token, err := authutil.GenerateJWTToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func NewLoginUsecase(userRepo repository.UserRepository) LoginUsecase {
	return &loginUsecaseImpl{
		userRepo: userRepo,
	}
}
