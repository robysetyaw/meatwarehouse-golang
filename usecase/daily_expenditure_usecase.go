package usecase

import (
	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type DailyExpenditureUseCase struct {
	dailyExpenditureRepository *repository.DailyExpenditureRepository
}

func NewDailyExpenditureUseCase(dailyExpenditureRepository *repository.DailyExpenditureRepository) *DailyExpenditureUseCase {
	return &DailyExpenditureUseCase{
		dailyExpenditureRepository: dailyExpenditureRepository,
	}
}

func (uc *DailyExpenditureUseCase) CreateDailyExpenditure(dailyExpenditure *model.DailyExpenditure) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum menyimpan pengeluaran harian ke dalam repository
	// ...

	err := uc.dailyExpenditureRepository.CreateDailyExpenditure(dailyExpenditure)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat membuat pengeluaran harian
		// ...
		return err
	}


	return nil
}

func (uc *DailyExpenditureUseCase) UpdateDailyExpenditure(dailyExpenditure *model.DailyExpenditure) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum mengupdate pengeluaran harian di dalam repository
	// ...

	err := uc.dailyExpenditureRepository.UpdateDailyExpenditure(dailyExpenditure)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengupdate pengeluaran harian
		// ...
		return err
	}

	return nil
}

func (uc *DailyExpenditureUseCase) DeleteDailyExpenditure(dailyExpenditureID string) error {
	// Lakukan validasi atau logika bisnis lainnya sebelum menghapus pengeluaran harian di dalam repository
	// ...

	err := uc.dailyExpenditureRepository.DeleteDailyExpenditure(dailyExpenditureID)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat menghapus pengeluaran harian
		// ...
		return err
	}

	return nil
}

func (uc *DailyExpenditureUseCase) GetDailyExpenditureByID(dailyExpenditureID string) (*model.DailyExpenditure, error) {
	dailyExpenditure, err := uc.dailyExpenditureRepository.GetDailyExpenditureByID(dailyExpenditureID)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengambil pengeluaran harian berdasarkan ID
		// ...
		return nil, err
	}

	return dailyExpenditure, nil
}

func (uc *DailyExpenditureUseCase) GetAllDailyExpenditures() ([]*model.DailyExpenditure, error) {
	dailyExpenditures, err := uc.dailyExpenditureRepository.GetAllDailyExpenditures()
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat mengambil daftar pengeluaran harian
		// ...
		return nil, err
	}

	return dailyExpenditures, nil
}
