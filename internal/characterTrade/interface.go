package charactertrade

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryCharacterTrade interface {
	BulkCreateCharacterTrade(characterTradeList []domain.CharacterTrade) error
	GetByTradeEventId(tradeEventId int) ([]domain.CharacterTrade, error)
	DeleteByTradeEventId(tradeEventId int) (error)
}

type ServiceCharacterTrade interface {
	BulkCreateCharacterTrade(characterTradeList []domain.CharacterTrade)  error
	GetByTradeEventId(tradeEventId int) ([]domain.CharacterTrade, error)
	DeleteByTradeEventId(tradeEventId int) (error)
}