package backgroundXproficiency

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repository BackgroundXProficiencyRepository
}

func NewBackgroundXProficiencyService(repository BackgroundXProficiencyRepository) BackgroundXProficiencyService {
	return &service{repository: repository}
}

func (s *service) Create(backgroundXProficiency domain.BackgroundXProficiency) (domain.BackgroundXProficiency, error) {
	newBackgroundXProficiency, err := s.repository.Create(backgroundXProficiency)
	if err != nil {
		return domain.BackgroundXProficiency{}, err
	}
	return newBackgroundXProficiency, nil
}

func (s *service) Delete(backgroundXProficiency domain.BackgroundXProficiency) error {
	return s.repository.Delete(backgroundXProficiency)
}
