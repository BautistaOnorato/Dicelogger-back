package dice_event

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repository DiceEventRepository
}

func NewDiceEventService(repository DiceEventRepository) DiceEventService {
	return &service{repository: repository}
}

// DeleteByProtagonistId implements DiceEventService.
func (s *service) DeleteByProtagonistId(id int) error {
	return s.repository.DeleteByProtagonistId(id)
}


// GetBySessionId implements DiceEventService.
func (s *service) GetBySessionId(sessionid int) ([]domain.DiceEvent, error) {
	return s.repository.GetBySessionId(sessionid)
}

func (s *service) Create(diceEvent domain.DiceEvent) (domain.DiceEvent, error) {
	return s.repository.Create(diceEvent)
}
func (s *service) GetAll() ([]domain.DiceEvent, error) {
	return s.repository.GetAll()
}
func (s *service) GetById(id int) (domain.DiceEvent, error) {
	return s.repository.GetById(id)
}
func (s *service) Update(diceEvent domain.DiceEvent, id int) (domain.DiceEvent, error) {
	return s.repository.Update(diceEvent, id)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
