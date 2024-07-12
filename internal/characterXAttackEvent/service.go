package characterxattackevent

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	characterXAttackEventRepository CharacterXAttackEventRepository
}

func NewCharacterXAttackEventService(characterXAttackEventRepository CharacterXAttackEventRepository) CharacterXAttackEventService {
	return &service{characterXAttackEventRepository: characterXAttackEventRepository}
}

func (s *service) GetAll() ([]domain.CharacterXAttackEvent, error) {
	characterXAttackEvents, err := s.characterXAttackEventRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return characterXAttackEvents, nil
}

func (s *service) GetById(id int) (domain.CharacterXAttackEvent, error) {
	characterXAttackEvent, err := s.characterXAttackEventRepository.GetById(id)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}

	return characterXAttackEvent, nil
}

func (s *service) GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error) {
	characterXAttackEvents, err := s.characterXAttackEventRepository.GetByCharacterId(characterId)
	if err != nil {
		return nil, err
	}

	return characterXAttackEvents, nil
}

func (s *service) GetByEventId(attackEventId int) ([]domain.CharacterXAttackEvent, error) {
	characterXAttackEvents, err := s.characterXAttackEventRepository.GetByEventId(attackEventId)
	if err != nil {
		return nil, err
	}

	return characterXAttackEvents, nil
}

func (s *service) Create(characterXAttackEvent dto.CharacterXAttackEventDto) (domain.CharacterXAttackEvent, error) {
	newCharacterXAttackEvent := domain.CharacterXAttackEvent{
		CharacterId: characterXAttackEvent.CharacterId,
		EventId: characterXAttackEvent.EventId,
		Dmg: characterXAttackEvent.Dmg,
		DmgRoll: characterXAttackEvent.DmgRoll,
		AttackResult: characterXAttackEvent.AttackResult,
		AttackRoll: characterXAttackEvent.AttackRoll,
		ArmorClass: characterXAttackEvent.ArmorClass,
	}

	createdCharacterXAttackEvent, err := s.characterXAttackEventRepository.Create(newCharacterXAttackEvent)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}

	return createdCharacterXAttackEvent, nil
}

func (s *service) Delete(id int) error {
	err := s.characterXAttackEventRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
