package usecase

import (
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type UserUseCase struct {
	userRepository *repository.UserRepository
}

func NewUserUseCase(userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserUseCase) CreateUser(user *model.User) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum menyimpan pengguna ke dalam repository
	// ...

	err := uc.userRepository.CreateUser(user)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat membuat pengguna
		// ...
		return err
	}

	return nil
}

func (uc *UserUseCase) UpdateUser(user *model.User) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum mengupdate pengguna di dalam repository
	// ...

	err := uc.userRepository.UpdateUser(user)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengupdate pengguna
		// ...
		return err
	}

	return nil
}

func (uc *UserUseCase) DeleteUser(userID string) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum menghapus pengguna di dalam repository
	// ...

	err := uc.userRepository.DeleteUser(userID)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat menghapus pengguna
		// ...
		return err
	}

	return nil
}

func (uc *UserUseCase) GetUserByID(userID string) (*model.User, error) {
	user, err := uc.userRepository.GetUserByID(userID)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengambil pengguna berdasarkan ID
		// ...
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetAllUsers() ([]*model.User, error) {
	users, err := uc.userRepository.GetAllUsers()
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengambil daftar pengguna
		// ...
		return nil, err
	}
	

	return users, nil
}
