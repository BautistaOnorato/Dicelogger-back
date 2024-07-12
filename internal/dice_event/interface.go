package dice_event

import "github.com/proyecto-dnd/backend/internal/domain"

type DiceEventRepository interface {
	Create(diceEvent domain.DiceEvent) (domain.DiceEvent, error)
	GetAll() ([]domain.DiceEvent, error)
	GetById(id int) (domain.DiceEvent, error)
	GetBySessionId(sessionid int) ([]domain.DiceEvent, error)
	Update(diceEvent domain.DiceEvent, id int) (domain.DiceEvent, error)
	Delete(id int) error
	DeleteByProtagonistId(id int) error
}

type DiceEventService interface {
	Create(diceEvent domain.DiceEvent) (domain.DiceEvent, error)
	GetAll() ([]domain.DiceEvent, error)
	GetById(id int) (domain.DiceEvent, error)
	GetBySessionId(sessionid int) ([]domain.DiceEvent, error)
	Update(diceEvent domain.DiceEvent, id int) (domain.DiceEvent, error)
	Delete(id int) error
	DeleteByProtagonistId(id int) error
}
