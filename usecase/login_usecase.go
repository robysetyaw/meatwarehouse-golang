package usecase

import (
	"fmt"
	"time"

	// "enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	// VerifyLogin(username, password string) (*model.User, error)
	Login(username, password string) (string, error)
}

type loginUseCase struct {
	userRepository repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository) LoginUseCase {
	return &loginUseCase{
		userRepository: userRepo,
	}
}


func (uc *loginUseCase) Login(username, password string) (string, error) {
	// Mengecek apakah pengguna dengan username tersebut ada di penyimpanan data
	user, err := uc.userRepository.GetByUsername(username)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Verifikasi password pengguna dengan menggunakan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	// Menghasilkan token JWT
	token, err := generateJWTToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func generateJWTToken(userID, username, role string) (string, error) {
	// Membuat claim JWT
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
	}

	// Membuat token JWT dengan menggunakan secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("secret-key") // Ganti dengan secret key Anda sendiri
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}