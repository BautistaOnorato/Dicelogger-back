package saving_throws

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repository SavingThrowsRepository
}

func NewSavingThrowsService(repository SavingThrowsRepository) SavingThrowsService {
	return &service{repository: repository}
}

func (s *service) Create(savingThrowDto dto.SavingThrowDto) (domain.SavingThrow, error)
func (s *service) GetAll() ([]domain.SavingThrow, error)
func (s *service) GetById(id int) (domain.SavingThrow, error)
func (s *service) Update(savingThrowDto dto.SavingThrowDto, id int) (domain.SavingThrow, error)
func (s *service) Delete(id int) error
