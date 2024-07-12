package saving_throws

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type SavingThrowsRepository interface {
	Create(savingThrowDto dto.SavingThrowDto) (domain.SavingThrow, error)
	GetAll() ([]domain.SavingThrow, error)
	GetById(id int) (domain.SavingThrow, error)
	Update(savingThrowDto dto.SavingThrowDto, id int) (domain.SavingThrow, error)
	Delete(id int) error
}

type SavingThrowsService interface {
	Create(savingThrowDto dto.SavingThrowDto) (domain.SavingThrow, error)
	GetAll() ([]domain.SavingThrow, error)
	GetById(id int) (domain.SavingThrow, error)
	Update(savingThrowDto dto.SavingThrowDto, id int) (domain.SavingThrow, error)
	Delete(id int) error
}
