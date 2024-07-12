package class

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	classRepository RepositoryCharacterClass
}

func NewClassService(classRepository RepositoryCharacterClass) ClassService {
	return &service{classRepository: classRepository}
}

func (s *service) Create(classDto dto.ClassDto) (domain.Class, error) {

	createdClass, err := s.classRepository.Create(classDto)
	if err != nil {
		return domain.Class{}, err
	}

	return createdClass, nil
}

func (s *service) GetAll() ([]domain.Class, error) {
	classes, err := s.classRepository.GetAll()
	if err != nil {
		return []domain.Class{}, err
	}
	return classes, err
}

func (s *service) GetById(id int) (domain.Class, error) {
	class, err := s.classRepository.GetById(id)
	if err != nil {
		return domain.Class{}, err
	}
	return class, err
}

func (s *service) Update(classDto dto.ClassDto, id int) (domain.Class, error) {
	updatedClass, err := s.classRepository.Update(classDto, id)
	if err != nil {
		return domain.Class{}, err
	}
	return updatedClass, err
}

func (s *service) Delete(id int) error {
	return s.classRepository.Delete(id)
}
