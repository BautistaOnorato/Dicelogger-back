package skill

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repo RepositorySkill
}

func NewServiceSkill(repo RepositorySkill) ServiceSkill {
	return &service{repo: repo}
}

// Create implements ServiceSkill.
func (s *service) Create(skill domain.Skill) (domain.Skill, error) {
	newSkill, err := s.repo.Create(skill)
	if err != nil {
		return domain.Skill{}, err
	}
	return newSkill, nil
}

// Delete implements ServiceSkill.
func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements ServiceSkill.
func (s *service) GetAll() ([]domain.Skill, error) {
	skills, err := s.repo.GetAll()
	if err != nil {
		return []domain.Skill{}, err
	}
	return skills, nil
}


// GetByCharacterId implements ServiceSkill.
func (s *service) GetByCharacterId(characterId int) ([]domain.Skill, error) {
	skills, err := s.repo.GetByCharacterId(characterId)
	if err != nil {
		return []domain.Skill{}, err
	}
	return skills, nil
}

// GetByClassId implements ServiceSkill.
func (s *service) GetByClassId(classId int) ([]domain.Skill, error) {
	skills, err := s.repo.GetByClassId(classId)
	if err != nil {
		return []domain.Skill{}, err
	}
	return skills, nil
}

// GetById implements ServiceSkill.
func (s *service) GetById(id int) (domain.Skill, error) {
	skill, err := s.repo.GetById(id)
	if err != nil {
		return domain.Skill{}, err
	}
	return skill, nil
}

// Update implements ServiceSkill.
func (s *service) Update(skill domain.Skill) (domain.Skill, error) {
	skill, err := s.repo.Update(skill)
	if err != nil {
		return domain.Skill{}, err
	}
	return skill, nil
}
