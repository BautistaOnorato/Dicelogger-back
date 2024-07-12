package classXspell

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	classXspellRepository RepositoryClassXSpell
}

func NewClassXSpelService(repository RepositoryClassXSpell) ServiceClassXSpell {
	return &service{classXspellRepository: repository}
}

func (s *service) Create(classXSpellData domain.ClassXSpell) (domain.ClassXSpell, error) {
	createdClassXSpell, err := s.classXspellRepository.Create(classXSpellData)
	if err != nil {
		return domain.ClassXSpell{}, err
	}
	return createdClassXSpell, nil
}

func (s *service) Delete(classXClassData domain.ClassXSpell) error {
	return s.classXspellRepository.Delete(classXClassData);	
}
