package backgroundxSkill

import "github.com/proyecto-dnd/backend/internal/domain"

type BackgroundXSkillsService struct {
	backgroundXSkillsRepo RepositoryBackgroundXSkills
}

func (s *BackgroundXSkillsService) CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error) {
	newBackgroundXSkills, err := s.backgroundXSkillsRepo.CreateBackgroundXSkills(data)
	if err != nil {
		return domain.BackgroundXSkills{}, err
	}
	return newBackgroundXSkills, nil
}

func (s *BackgroundXSkillsService) DeleteBackgroundXSkills(id int) error {
	err := s.backgroundXSkillsRepo.DeleteBackgroundXSkills(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BackgroundXSkillsService) DeleteByBackgroundId(id int) error {
	err := s.backgroundXSkillsRepo.DeleteByBackgroundId(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BackgroundXSkillsService) GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error) {
	backgroundXSkills, err := s.backgroundXSkillsRepo.GetAllBackgroundXSkills()
	if err != nil {
		return []domain.BackgroundXSkills{}, err
	}
	return backgroundXSkills, nil
}

func (s *BackgroundXSkillsService) GetByBackgroundId(id int) ([]domain.BackgroundXSkills, error) {
	backgroundXSkills, err := s.backgroundXSkillsRepo.GetByBackgroundId(id)
	if err != nil {
		return []domain.BackgroundXSkills{}, err
	}
	return backgroundXSkills, nil
}

func (s *BackgroundXSkillsService) GetByIdBackgroundXSkills(id int) (domain.BackgroundXSkills, error) {
	backgroundXSkills, err := s.backgroundXSkillsRepo.GetByIdBackgroundXSkills(id)
	if err != nil {
		return domain.BackgroundXSkills{}, err
	}
	return backgroundXSkills, nil
}

func (s *BackgroundXSkillsService) UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error) {
	updatedBackgroundXSkills, err := s.backgroundXSkillsRepo.UpdateBackgroundXSkills(data)
	if err != nil {
		return domain.BackgroundXSkills{}, err
	}
	return updatedBackgroundXSkills, nil
}

func NewBackgroundXSkillsService(repo RepositoryBackgroundXSkills) ServiceBackgroundXSkills {
	return &BackgroundXSkillsService{backgroundXSkillsRepo: repo}
}
